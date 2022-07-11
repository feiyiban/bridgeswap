package connection

import (
	"fmt"
	"math"
	"math/big"
	"sync"
	"time"

	// "crosschainbridge/controller/crypto/secp256k1"
	"bridgeswap/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	trongrpc "bridgeswap/chains/tron/pkg/client"
	troncommon "bridgeswap/chains/tron/pkg/common"
	tronapi "bridgeswap/chains/tron/pkg/proto/api"
)

var BlockRetryInterval = time.Second * 5

type Connection struct {
	withTLS  bool
	endpoint string
	conn     *trongrpc.GrpcClient

	// kp            *secp256k1.Keypair
	gasLimit      *big.Int
	maxGasPrice   *big.Int
	minGasPrice   *big.Int
	gasMultiplier *big.Float
	egsApiKey     string
	egsSpeed      string

	optsLock sync.Mutex
	log      logger.Logger
	stop     chan int // All routines should exit when this channel is closed
}

// NewConnection returns an uninitialized connection, must call Connection.Connect() before using.
func NewConnection(endpoint string, withTLS bool, log logger.Logger, gasLimit, maxGasPrice, minGasPrice *big.Int, gasMultiplier *big.Float, gsnApiKey, gsnSpeed string) *Connection {
	return &Connection{
		endpoint: endpoint,
		withTLS:  withTLS,

		gasLimit:      gasLimit,
		maxGasPrice:   maxGasPrice,
		minGasPrice:   minGasPrice,
		gasMultiplier: gasMultiplier,
		egsApiKey:     gsnApiKey,
		egsSpeed:      gsnSpeed,
		log:           log,
		stop:          make(chan int),
	}
}

func (c *Connection) Connect() error {
	c.log.Info("Connecting to tron chain...", "url", c.endpoint)
	grpcClient := trongrpc.NewGrpcClient(c.endpoint)

	// load grpc options
	opts := make([]grpc.DialOption, 0)
	if c.withTLS {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	if err := grpcClient.Start(opts...); err != nil {
		return err
	}

	c.conn = grpcClient

	return nil
}

func (c *Connection) Client() *trongrpc.GrpcClient {
	return c.conn
}

func (c *Connection) SelfChainId(contractAddress string) (*big.Int, error) {
	result, err := c.Client().TriggerConstantContract("", contractAddress, "selfChainId()", "")
	if err != nil {
		return big.NewInt(0), err
	}

	data := troncommon.BytesToHexString(result.GetConstantResult()[0])

	c.log.Info("jsonData", data)

	return c.Client().ParseTRC20NumericProperty(data)

}

func (c *Connection) DepositOut(from, contractAddress, param string, feeLimit int64, tAmount float64, tTokenID string, tTokenAmount int64) (*tronapi.TransactionExtention, error) {
	valueInt := int64(0)
	if tAmount > 0 {
		valueInt = int64(tAmount * math.Pow10(6))
	}

	tokenInt := int64(0)
	if tTokenAmount > 0 {
		// get token info
		info, err := c.Client().GetAssetIssueByID(tTokenID)
		if err != nil {
			return nil, err
		}
		tokenInt = int64(tAmount * math.Pow10(int(info.Precision)))
	}

	tx, err := c.Client().TriggerContract(from, contractAddress, "depositOut(address,address,address,uint256)",
		param,
		feeLimit,
		valueInt,
		tTokenID,
		tokenInt,
	)
	if err != nil {
		return nil, err
	}

	return tx, nil

}

// LatestBlock returns the latest block from the current chain
func (c *Connection) LatestBlock() (*big.Int, error) {
	header, err := c.conn.GetNowBlock()
	if err != nil {
		return nil, err
	}

	if header.BlockHeader == nil || header.BlockHeader.RawData == nil {
		return nil, fmt.Errorf("no found latest block")
	}

	height := header.BlockHeader.RawData.Number

	return big.NewInt(height), nil
}

// LatestBlock returns the latest block from the current chain
// func (c *Connection) LatestBlock() (*big.Int, error) {
// 	header, err := c.Client().HeaderByNumber(context.Background(), nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return header.Number, nil
// }

// func (c *Connection) PendingNonceAt(ctx context.Context, account common.Address) (*big.Int, error) {
// 	return c.conn.PendingBalanceAt(ctx, account)
// }

// // newTransactOpts builds the TransactOpts for the connection's keypair.
// func (c *Connection) newTransactOpts(value, gasLimit, gasPrice *big.Int) (*bind.TransactOpts, error) {
// 	privateKey := c.kp.PrivateKey()
// 	address := ethcrypto.PubkeyToAddress(privateKey.PublicKey)

// 	c.log.Info("newTransactopts", "address", address.Hex())
// 	// nonce, err := c.conn.PendingNonceAt(context.Background(), address)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	id, err := c.conn.ChainID(context.Background())
// 	if err != nil {
// 		return nil, err
// 	}

// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = value
// 	auth.GasLimit = uint64(gasLimit.Int64())
// 	auth.GasPrice = gasPrice
// 	auth.Context = context.Background()

// 	return auth, nil
// }

// // EnsureHasBytecode asserts if contract code exists at the specified address
// func (c *Connection) EnsureHasBytecode(addr common.Address) error {
// 	code, err := c.conn.CodeAt(context.Background(), addr, nil)
// 	if err != nil {
// 		return err

// 	}

// 	if len(code) == 0 {
// 		return fmt.Errorf("no bytecode found at %s", addr.Hex())
// 	}
// 	return nil
// }

func multiplyGasPrice(gasEstimate *big.Int, gasMultiplier *big.Float) *big.Int {

	gasEstimateFloat := new(big.Float).SetInt(gasEstimate)

	result := gasEstimateFloat.Mul(gasEstimateFloat, gasMultiplier)

	gasPrice := new(big.Int)

	result.Int(gasPrice)

	return gasPrice
}

// LockAndUpdateOpts acquires a lock on the opts before updating the nonce
// and gas price.

func (c *Connection) UnlockOpts() {
	c.optsLock.Unlock()
}

// Close terminates the client connection and stops any running routines
func (c *Connection) Close() {
	if c.conn != nil {
	}
	c.conn.Stop()
	close(c.stop)
}
