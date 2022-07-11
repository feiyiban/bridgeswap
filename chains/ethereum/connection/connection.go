package connection

import (
	"bridgeswap/chains/ethereum/crypto/secp256k1"
	"bridgeswap/logger"
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Connection struct {
	endpoint      string
	http          bool
	kp            *secp256k1.Keypair
	gasLimit      *big.Int
	maxGasPrice   *big.Int
	minGasPrice   *big.Int
	gasMultiplier *big.Float

	client *ethclient.Client
	// signer    ethtypes.Signer
	opts     *bind.TransactOpts
	callOpts *bind.CallOpts
	nonce    uint64
	optsLock sync.Mutex
	log      logger.Logger
	stop     chan int // All routines should exit when this channel is closed
}

// NewConnection returns an uninitialized connection, must call Connection.Connect() before using.
func NewConnection(endpoint string, http bool, kp *secp256k1.Keypair, log logger.Logger, gasLimit, maxGasPrice, minGasPrice *big.Int, gasMultiplier *big.Float) *Connection {
	return &Connection{
		endpoint:      endpoint,
		http:          http,
		kp:            kp,
		gasLimit:      gasLimit,
		maxGasPrice:   maxGasPrice,
		minGasPrice:   minGasPrice,
		gasMultiplier: gasMultiplier,
		log:           log,
		stop:          make(chan int),
	}
}

// Connect starts the ethereum WS connection
func (conn *Connection) Connect() error {
	conn.log.Info("Connecting to ethereum chain...", "url", conn.endpoint)
	var rpcClient *rpc.Client
	var err error
	// Start http or ws client
	if conn.http {
		rpcClient, err = rpc.DialHTTP(conn.endpoint)
	} else {
		rpcClient, err = rpc.DialContext(context.Background(), conn.endpoint)
	}
	if err != nil {

		return err
	}

	conn.client = ethclient.NewClient(rpcClient)

	return nil
}

func (conn *Connection) Client() *ethclient.Client {
	return conn.client
}

// LatestBlock returns the latest block from the current chain
func (conn *Connection) LatestBlock() (*big.Int, error) {
	header, err := conn.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

// Close terminates the client connection and stops any running routines
func (conn *Connection) Close() {
	if conn.client != nil {
	}
	conn.client.Close()
	close(conn.stop)
}
