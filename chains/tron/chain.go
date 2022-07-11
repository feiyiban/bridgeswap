package tron

import (
	// metrics "crosschainbridge/controller/metrics/types"
	"fmt"
	"math/big"

	"bridgeswap/controller/core"
	// "bridgeswap/controller/crypto/secp256k1"

	// metrics "crosschainbridge/controller/metrics/types"

	// "crosschainbridge/chains/tron/addressexchange"
	"bridgeswap/chains/tron/connection"

	"bridgeswap/logger"

	tronGrpc "bridgeswap/chains/tron/pkg/client"
	tronapi "bridgeswap/chains/tron/pkg/proto/api"
)

type Connection interface {
	Connect() error
	Client() *tronGrpc.GrpcClient
	LatestBlock() (*big.Int, error)
	DepositOut(from, contractAddress, param string, feeLimit int64, tAmount float64, tTokenID string, tTokenAmount int64) (*tronapi.TransactionExtention, error)
	Close()
}

type Chain struct {
	cfg      *core.ChainConfig // The config of the chain
	conn     Connection        // THe chains connection
	listener *listener         // The listener of this chain
	writer   *writer           // The writer of the chain
	stop     chan<- int
}

// checkBlockstore queries the blockstore for the latest known block. If the latest block is
// greater than cfg.startBlock, then cfg.startBlock is replaced with the latest known block.
// func setupBlockstore(cfg *Config, kp *secp256k1.Keypair) (*blockstore.Blockstore, error) {
// 	bs, err := blockstore.NewBlockstore(cfg.blockstorePath, cfg.id, kp.Address())
// 	if err != nil {
// 		return nil, err
// 	}

// 	if !cfg.freshStart {
// 		latestBlock, err := bs.TryLoadLatestBlock()
// 		if err != nil {
// 			return nil, err
// 		}

// 		if latestBlock.Cmp(cfg.startBlock) == 1 {
// 			cfg.startBlock = latestBlock
// 		}
// 	}

// 	return bs, nil
// }

func InitializeChain(chainCfg *core.ChainConfig, log logger.Logger, sysErr chan<- error) (*Chain, error) {
	cfg, err := parseChainConfig(chainCfg)
	if err != nil {
		return nil, err
	}

	// bs, err := setupBlockstore(cfg, kp)
	// if err != nil {
	// 	log.Debug("setupBlockstore:", err.Error())
	// 	return nil, err
	// }

	stop := make(chan int)
	conn := connection.NewConnection(cfg.endpoint, cfg.http, log, cfg.gasLimit, cfg.maxGasPrice, cfg.minGasPrice, cfg.gasMultiplier, cfg.egsApiKey, cfg.egsSpeed)
	err = conn.Connect()
	if err != nil {
		log.Debug("Connect:", err.Error())
		return nil, err
	}

	chainId, err := conn.SelfChainId(cfg.bridgeContract.String())
	if err != nil {
		log.Debug("bridge contract selfchainid:", err.Error())
		return nil, err
	}

	if chainId.Cmp(new(big.Int).SetUint64(uint64(chainCfg.ID))) != 0 {
		return nil, fmt.Errorf("chainId (%d) and configuration chainId (%d) do not match", chainId, chainCfg.ID)
	}

	if chainCfg.LatestBlock {
		curr, err := conn.LatestBlock()
		if err != nil {
			return nil, err
		}
		cfg.startBlock = curr
	}

	listener := NewListener(conn, cfg, log, stop, sysErr)
	// listener.setContracts(bridgeContract)

	writer := NewWriter(conn, cfg, log, stop, sysErr)
	writer.setContract(cfg.bridgeContract.String())

	return &Chain{
		cfg:      chainCfg,
		conn:     conn,
		writer:   writer,
		listener: listener,
		stop:     stop,
	}, nil
}

func (c *Chain) SetRouter(r *core.Router) {
	c.writer.log.Info("SetRouter")
	r.Listen(c.cfg.ID, c.writer)
	c.listener.setRouter(r)
}

func (c *Chain) Start() error {
	c.writer.log.Debug("started chain...")
	err := c.listener.start()
	if err != nil {
		return err
	}

	err = c.writer.start()
	if err != nil {
		return err
	}

	c.writer.log.Debug("Successfully started chain")
	return nil
}

func (c *Chain) ID() uint8 {
	return c.cfg.ID
}

func (c *Chain) Name() string {
	return c.cfg.Name
}

// func (c *Chain) LatestBlock() metrics.LatestBlock {
// 	return c.listener.latestBlock
// }

// Stop signals to any running routines to exit
func (c *Chain) Stop() {
	close(c.stop)
	if c.conn != nil {
		c.conn.Close()
	}
}
