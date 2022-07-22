package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"bridgeswap/bindings/eth/bridgev1"
	"bridgeswap/blockstore"
	"bridgeswap/chains/ethereum/connection"
	"bridgeswap/controller/core"
	"bridgeswap/logger"
	"bridgeswap/sdk/ethereum/crypto/secp256k1"
	"bridgeswap/sdk/ethereum/keystore"
	"math/big"
)

type Connection interface {
	Connect() error
	Keypair() *secp256k1.Keypair
	Client() *ethclient.Client
	EnsureHasBytecode(addr common.Address) error
	LatestBlock() (*big.Int, error)
	Close()
}

type Chain struct {
	cfg      *core.ChainConfig // The config of the chain
	conn     Connection        // THe chains connection
	listener *listener         // The listener of this chain
	writer   *writer           // The writer of the chain
	stop     chan<- int
}

func InitializeChain(chainCfg *core.ChainConfig, sysErr chan<- error, log logger.Logger) (*Chain, error) {
	cfg, err := parseChainConfig(chainCfg)
	if err != nil {
		return nil, err
	}

	cryptoKeyPair, err := keystore.KeypairFromAddress(cfg.from, keystore.EthChain, cfg.keystorePath)
	if err != nil {
		return nil, err
	}
	secp256k1KeyPair, ok := cryptoKeyPair.(*secp256k1.Keypair)
	if !ok {
		return nil, err
	}

	blockStore, err := blockstore.NewBlockstore(cfg.blockstorePath, cfg.id, cfg.from)
	if err != nil {
		return nil, err
	}
	if !cfg.bFreshStart {
		latestBlock, err := blockStore.TryLoadLatestBlock()
		if err != nil {
			return nil, err
		}

		if latestBlock.Cmp(cfg.startBlock) == 1 {
			cfg.startBlock = latestBlock
		}
	}

	conn := connection.NewConnection(cfg.http, cfg.endpoint, secp256k1KeyPair, cfg.gasLimit, cfg.maxGasPrice, cfg.minGasPrice, log)
	err = conn.Connect()
	if err != nil {
		return nil, err
	}

	err = conn.EnsureHasBytecode(common.HexToAddress(cfg.bridgeContract))
	if err != nil {
		return nil, err
	}

	if cfg.erc20Contract != "" {
		err = conn.EnsureHasBytecode(common.HexToAddress(cfg.erc20Contract))
		if err != nil {
			return nil, err
		}
	}

	bridge, err := bridgev1.NewBridgev1(common.HexToAddress(cfg.bridgeContract), conn.Client())
	if err != nil {
		log.Debug("error happened when call NewBridgev1", "error", err)
		return nil, err
	}

	if cfg.bLatestBlock {
		curr, err := conn.LatestBlock()
		if err != nil {
			return nil, err
		}
		cfg.startBlock = curr
	}

	stop := make(chan int)
	listener := NewListener(cfg, conn, blockStore, sysErr, stop, log)
	listener.setContracts(bridge)

	writer := NewWriter(cfg, conn, sysErr, stop, log)
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
