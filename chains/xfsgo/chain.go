package xfsgo

import (
	"bridgeswap/blockstore"
	"bridgeswap/chains/xfsgo/connection"
	"bridgeswap/chains/xfsgo/types"
	"bridgeswap/controller/core"
	"bridgeswap/logger"
	"bridgeswap/sdk/xfs/rpcclient"
	"math/big"
)

type Connection interface {
	Connect() error
	Client() *rpcclient.Client
	LatestBlock() (*big.Int, error)
	SendRawTransaction(rawTx string) (string, error)
	SignedTx(args types.StringRawTransaction) (string, error)
	GetLogs(args types.GetLogsRequest) (*[]*types.EventLogResp, error)
}

type Chain struct {
	cfg      *core.ChainConfig // The config of the chain
	conn     Connection        // THe chains connection
	listener *listener         // The listener of this chain
	writer   *writer           // The writer of the chain
	stop     chan<- int
}

// checkBlockstore queries the blockstore for the latest known block. If the latest block is
// greater than startBlock, then the latest block is returned, otherwise startBlock is.
func checkBlockstore(bs *blockstore.Blockstore, startBlock uint64) (uint64, error) {
	latestBlock, err := bs.TryLoadLatestBlock()
	if err != nil {
		return 0, err
	}

	if latestBlock.Uint64() > startBlock {
		return latestBlock.Uint64(), nil
	} else {
		return startBlock, nil
	}
}

func InitializeChain(cfg *core.ChainConfig, log logger.Logger, sysErr chan<- error) (*Chain, error) {
	xfsCfg, err := parseChainConfig(cfg)
	if err != nil {
		return nil, err
	}

	bs, err := blockstore.NewBlockstore(cfg.BlockstorePath, cfg.ID, cfg.From)
	if err != nil {
		return nil, err
	}

	stop := make(chan int)
	// Setup connection
	conn := connection.NewConnection(xfsCfg.endpoint, xfsCfg.name, nil, log, stop, sysErr)
	err = conn.Connect()
	if err != nil {
		log.Debug("Connection", cfg.ID, err)
		return nil, err
	}

	ue := parseUseExtended(cfg)

	// Setup listener & writer
	l := NewListener(conn, xfsCfg, log, bs, stop, sysErr)
	w := NewWriter(conn, xfsCfg, log, sysErr, ue)
	return &Chain{
		cfg:      cfg,
		conn:     conn,
		listener: l,
		writer:   w,
		stop:     stop,
	}, nil
}

func (chain *Chain) SetRouter(r *core.Router) {
	chain.writer.log.Info("SetRouter")
	r.Listen(chain.cfg.ID, chain.writer)
	chain.listener.setRouter(r)
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
		// c.conn.Close()
	}
}
