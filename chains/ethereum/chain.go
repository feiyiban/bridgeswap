package ethereum

import (
	"bridgeswap/bindings/eth/bridgev1"
	"bridgeswap/blockstore"
	"bridgeswap/chains/ethereum/connection"
	"bridgeswap/chains/ethereum/crypto/secp256k1"
	"bridgeswap/chains/ethereum/keystore"
	"bridgeswap/controller/core"
	"bridgeswap/logger"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Connection interface {
	Connect() error
	Keypair() *secp256k1.Keypair
	Client() *ethclient.Client
	LatestBlock() (*big.Int, error)
	// Opts() *bind.TransactOpts
	// CallOpts() *bind.CallOpts
	// LockAndUpdateOpts() error
	// UnlockOpts()

	// EnsureHasBytecode(address common.Address) error
	// LatestBlock() (*big.Int, error)
	// WaitForBlock(block *big.Int, delay *big.Int) error
	Close()
}

type Chain struct {
	cfg  *core.ChainConfig // The config of the chain
	conn Connection        // THe chains connection

	listener *listener // The listener of this chain
	writer   *writer   // The writer of the chain

	stop chan<- int
}

// checkBlockstore queries the blockstore for the latest known block. If the latest block is
// greater than cfg.startBlock, then cfg.startBlock is replaced with the latest known block.
func setupBlockstore(cfg *Config, addr string) (*blockstore.Blockstore, error) {
	bs, err := blockstore.NewBlockstore(cfg.blockstorePath, cfg.id, addr)
	if err != nil {
		return nil, err
	}

	if !cfg.freshStart {
		latestBlock, err := bs.TryLoadLatestBlock()
		if err != nil {
			return nil, err
		}

		if latestBlock.Cmp(cfg.startBlock) == 1 {
			cfg.startBlock = latestBlock
		}
	}

	return bs, nil
}

func InitializeChain(chainCfg *core.ChainConfig, log logger.Logger, sysErr chan<- error) (*Chain, error) {
	cfg, err := parseChainConfig(chainCfg)
	if err != nil {
		return nil, err
	}

	kpI, err := keystore.KeypairFromAddress(cfg.from, keystore.EthChain, cfg.keystorePath)
	if err != nil {
		return nil, err
	}

	kp, ok := kpI.(*secp256k1.Keypair)
	if !ok {
		return nil, fmt.Errorf("keystore %s", "Get Keypair err")
	}

	bs, err := setupBlockstore(cfg, cfg.from)
	if err != nil {
		log.Debug("setupBlockstore:", err.Error())
		return nil, err
	}

	conn := connection.NewConnection(cfg.endpoint, cfg.http, kp, log, cfg.gasLimit, cfg.maxGasPrice, cfg.minGasPrice, cfg.gasMultiplier)
	err = conn.Connect()
	if err != nil {
		log.Debug("Connect:", err.Error())
		return nil, err
	}

	bridge, err := bridgev1.NewBridgev1(cfg.bridgeContract, conn.Client())
	if err != nil {
		log.Debug("Get Bridge Obj Err by bridgeContract", err.Error())
		return nil, err
	}

	stop := make(chan int)

	listener := NewListener(conn, cfg, log, bs, stop, sysErr)
	listener.setContracts(bridge)

	writer := NewWriter(conn, cfg, log, stop, sysErr)
	writer.setContract(bridge)

	return &Chain{
		cfg:      chainCfg,
		conn:     conn,
		writer:   writer,
		listener: listener,
		stop:     stop,
	}, nil

}

func (chain *Chain) SetRouter(router *core.Router) {
	chain.writer.log.Info("SetRouter")
	router.Listen(chain.cfg.ID, chain.writer)
	chain.listener.setRouter(router)
}

func (chain *Chain) Start() error {
	chain.writer.log.Debug("started chain...")
	err := chain.listener.start()
	if err != nil {
		return err
	}

	err = chain.writer.start()
	if err != nil {
		return err
	}

	chain.writer.log.Debug("Successfully started chain")

	return nil
}

func (chain *Chain) ID() uint8 {
	return chain.cfg.ID
}

func (chain *Chain) Name() string {
	return chain.cfg.Name
}

func (chain *Chain) Stop() {
	close(chain.stop)
	if chain.conn != nil {
		chain.conn.Close()
	}
}
