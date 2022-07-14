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
	"bridgeswap/chains/tron/pkg/keystore"

	trontransaction "bridgeswap/chains/tron/pkg/client/transaction"
	tronstore "bridgeswap/chains/tron/pkg/store"
)

var BlockRetryInterval = time.Second * 5

type Connection struct {
	withTLS       bool
	endpoint      string
	conn          *trongrpc.GrpcClient
	key           *keystore.KeyStore
	senderAcct    *keystore.Account
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
func NewConnection(endpoint string, withTLS bool, log logger.Logger, key *keystore.KeyStore, senderAcct *keystore.Account, gasLimit, maxGasPrice, minGasPrice *big.Int, gasMultiplier *big.Float, gsnApiKey, gsnSpeed string) *Connection {
	return &Connection{
		endpoint:      endpoint,
		withTLS:       withTLS,
		log:           log,
		key:           key,
		senderAcct:    senderAcct,
		gasLimit:      gasLimit,
		maxGasPrice:   maxGasPrice,
		minGasPrice:   minGasPrice,
		gasMultiplier: gasMultiplier,
		egsApiKey:     gsnApiKey,
		egsSpeed:      gsnSpeed,

		stop: make(chan int),
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

	c.log.Info("tron", "chainid", data)

	return c.Client().ParseTRC20NumericProperty(data)

}

func opts(ctlr *trontransaction.Controller) {
}

func (c *Connection) TransferIn(from, contractAddress, param string, feeLimit int64, tAmount float64, tTokenID string, tTokenAmount int64) error {
	valueInt := int64(0)
	if tAmount > 0 {
		valueInt = int64(tAmount * math.Pow10(6))
	}

	tokenInt := int64(0)
	if tTokenAmount > 0 {
		// get token info
		info, err := c.Client().GetAssetIssueByID(tTokenID)
		if err != nil {
			return err
		}
		tokenInt = int64(tAmount * math.Pow10(int(info.Precision)))
	}

	tx, err := c.Client().TriggerContract(from, contractAddress, "transferIn(address,address,uint256,uint256,uint256)",
		param,
		feeLimit,
		valueInt,
		tTokenID,
		tokenInt,
	)
	if err != nil {
		return err
	}

	c.log.Info("tron--->bridgeContract--->TransferIn", "tx", tx)

	ks, acct, err := tronstore.UnlockedKeystore(from, "")
	if err != nil {
		return err
	}
	ctrlr := trontransaction.NewController(c.Client(), ks, acct, tx.Transaction, opts)
	if err = ctrlr.ExecuteTransaction(); err != nil {
		c.log.Error("ExecuteTransaction", "TransferIn", err)
		return err
	}

	return nil
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
