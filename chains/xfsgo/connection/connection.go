package connection

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"bridgeswap/chains/xfsgo/pkg/types"
	"bridgeswap/chains/xfsgo/pkg/xfsclient"
	"bridgeswap/chains/xfsgo/pkg/xfsrpc"

	"bridgeswap/logger"
)

type Connection struct {
	endpoint string
	http     bool
	conn     *xfsclient.Client
	log      logger.Logger
	name     string // Chain name
	key      *ecdsa.PrivateKey
	stop     <-chan int   // Signals system shutdown, should be observed in all selects and loops
	sysErr   chan<- error // Propagates fatal errors to core
}

func NewConnection(endpoint string, http bool, name string, key *ecdsa.PrivateKey, log logger.Logger, stop <-chan int, sysErr chan<- error) *Connection {
	return &Connection{endpoint: endpoint, http: http, name: name, key: key, log: log, stop: stop, sysErr: sysErr}
}

func (c *Connection) Connect() error {
	c.log.Info("Connecting to xfsgo chain...", "url", c.endpoint)
	var rpcClient *xfsrpc.Client
	var err error
	// Start http or ws client
	if c.http {
		rpcClient, err = xfsrpc.DialHTTP(c.endpoint)
	} else {
		rpcClient, err = xfsrpc.DialContext(context.Background(), c.endpoint)
	}
	if err != nil {
		return err
	}
	c.conn = xfsclient.NewClient(rpcClient)

	return nil
}

// LatestBlock returns the latest block from the current chain
func (c *Connection) LatestBlock() (*big.Int, error) {
	header, err := c.conn.GetHead(context.Background())
	if err != nil {
		return nil, err
	}
	return header, nil
}

func (c *Connection) GetLogs(args types.GetLogsRequest) (*[]*types.EventLogResp, error) {
	eventLog, err := c.conn.GetLogs(context.Background(), args)
	if err != nil {
		return nil, err
	}
	return eventLog, nil
}

func (c *Connection) SignedTx(args types.StringRawTransaction) (*string, error) {
	eventLog, err := c.conn.SignedTx(context.Background(), args)
	if err != nil {
		return nil, err
	}
	return eventLog, nil
}

func (c *Connection) SendRawTransaction(rawTx string) (string, error) {
	txHash, err := c.conn.SendRawTransaction(context.Background(), rawTx)
	if err != nil {
		return "", err
	}
	return txHash, nil
}

func (c *Connection) checkChainId(expected uint8) error {
	// var actual msg.ChainId
	// err := c.getConst(BridgePalletName, "ChainIdentity", &actual)
	// if err != nil {
	// 	return err
	// }

	// if actual != expected {
	// 	return fmt.Errorf("ChainID is incorrect, Expected chainId: %d, got chainId: %d", expected, actual)
	// }

	return nil
}

func (c *Connection) Client() *xfsclient.Client {
	return c.Client()
}

func (c *Connection) TransferIn(token, contractAddress, param string, feeLimit int64, tAmount float64, tTokenID string, tTokenAmount int64) error {
	return nil
}
