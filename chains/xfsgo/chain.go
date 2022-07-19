/*
The substrate package contains the logic for interacting with substrate chains.
The current supported transfer types are Fungible, Nonfungible, and generic.

There are 3 major components: the connection, the listener, and the writer.

Connection

The Connection handles connecting to the substrate client, and submitting transactions to the client.
It also handles state queries. The connection is shared by the writer and listener.

Listener

The substrate listener polls blocks and parses the associated events for the three transfer types. It then forwards these into the router.

Writer

As the writer receives messages from the router, it constructs proposals. If a proposal is still active, the writer will attempt to vote on it. Resource IDs are resolved to method name on-chain, which are then used in the proposals when constructing the resulting Call struct.

*/
package xfsgo

import (
	"bridgeswap/blockstore"
	"bridgeswap/chains/xfsgo/connection"
	"bridgeswap/chains/xfsgo/pkg/types"
	"bridgeswap/chains/xfsgo/pkg/xfsclient"
	"bridgeswap/controller/core"
	"bridgeswap/logger"
	"math/big"
)

type Connection interface {
	Connect() error
	Client() *xfsclient.Client
	LatestBlock() (*big.Int, error)
	TransferIn(token, contractAddress, param string, feeLimit int64, tAmount float64, tTokenID string, tTokenAmount int64) error
	SendRawTransaction(rawTx string) (string, error)
	SignedTx(args types.StringRawTransaction) (*string, error)
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
	conn := connection.NewConnection(xfsCfg.endpoint, xfsCfg.http, xfsCfg.name, nil, log, stop, sysErr)
	err = conn.Connect()
	if err != nil {
		log.Debug("Connection", cfg.ID, err)
		return nil, err
	}

	ue := parseUseExtended(cfg)

	// Setup listener & writer
	l := NewListener(conn, xfsCfg, log, bs, stop, sysErr)
	w := NewWriter(conn, log, sysErr, ue)
	return &Chain{
		cfg:      cfg,
		conn:     conn,
		listener: l,
		writer:   w,
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

func (c *Chain) Stop() {
	close(c.stop)
	if c.conn != nil {
		// c.conn.Close()
	}
}
