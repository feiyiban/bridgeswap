package connection

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"bridgeswap/logger"
	"bridgeswap/sdk/ethereum/crypto/secp256k1"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type Connection struct {
	http        bool
	endpoint    string
	keypair     *secp256k1.Keypair
	gasLimit    *big.Int
	maxGasPrice *big.Int
	minGasPrice *big.Int
	client      *ethclient.Client
	opts        *bind.TransactOpts
	callOpts    *bind.CallOpts
	nonce       uint64
	optsLock    sync.Mutex
	stop        chan int // All routines should exit when this channel is closed
	log         logger.Logger
}

// NewConnection returns an uninitialized connection, must call Connection.Connect() before using.
func NewConnection(http bool, endpoint string, keypair *secp256k1.Keypair, gasLimit, maxGasPrice, minGasPrice *big.Int, log logger.Logger) *Connection {
	return &Connection{
		http:        http,
		endpoint:    endpoint,
		keypair:     keypair,
		gasLimit:    gasLimit,
		maxGasPrice: maxGasPrice,
		minGasPrice: minGasPrice,
		stop:        make(chan int),
		log:         log,
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

func (conn *Connection) Keypair() *secp256k1.Keypair {
	return conn.keypair
}

// LatestBlock returns the latest block from the current chain
func (conn *Connection) LatestBlock() (*big.Int, error) {
	header, err := conn.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	return header.Number, nil
}

// EnsureHasBytecode asserts if contract code exists at the specified address
func (conn *Connection) EnsureHasBytecode(addr common.Address) error {
	code, err := conn.client.CodeAt(context.Background(), addr, nil)
	if err != nil {
		return err
	}

	if len(code) == 0 {
		return fmt.Errorf("no bytecode found at %s", addr.Hex())
	}
	return nil
}

// Close terminates the client connection and stops any running routines
func (conn *Connection) Close() {
	if conn.client != nil {
	}
	conn.client.Close()
	close(conn.stop)
}
