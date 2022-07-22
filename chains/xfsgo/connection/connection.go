package connection

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strconv"

	"bridgeswap/chains/xfsgo/types"
	"bridgeswap/sdk/xfsgo/common"
	"bridgeswap/sdk/xfsgo/rpcclient"

	"bridgeswap/logger"
)

type Connection struct {
	endpoint string            // client url
	cli      *rpcclient.Client // client object
	name     string            // Chain name
	key      *ecdsa.PrivateKey // External account private key
	stop     <-chan int        // Signals system shutdown, should be observed in all selects and loops
	sysErr   chan<- error      // Propagates fatal errors to core
	log      logger.Logger
}

func NewConnection(endpoint string, name string, key *ecdsa.PrivateKey, log logger.Logger, stop <-chan int, sysErr chan<- error) *Connection {
	return &Connection{endpoint: endpoint, name: name, key: key, log: log, stop: stop, sysErr: sysErr}
}

func (conn *Connection) Connect() error {
	conn.log.Info("Connecting to xfsgo chain...", "url", conn.endpoint)

	conn.cli = rpcclient.NewClient(conn.endpoint, "180")

	return nil
}

// LatestBlock returns the latest block from the current chain
func (conn *Connection) LatestBlock() (*big.Int, error) {
	var block map[string]interface{}
	err := conn.cli.CallMethod(1, "Chain.GetHead", nil, &block)
	if err != nil {
		return big.NewInt(int64(0)), err
	}
	if err != nil {
		return nil, err
	}

	height := block["height"].(float64)

	bigHeight := new(big.Float).SetFloat64(height)

	int64Height, _ := bigHeight.Int64()

	return big.NewInt(int64Height), err
}

func (conn *Connection) GetLogs(args types.GetLogsRequest) (*[]*types.EventLogResp, error) {
	resp := make([]*types.EventLogResp, 0)

	err := conn.cli.CallMethod(1, "Chain.GetLogs", args, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (conn *Connection) SignedTx(tx types.StringRawTransaction) (string, error) {
	if tx.Version == "" {
		tx.Version = "0"
	}
	if tx.Value == "" {
		tx.Value = "0"
	}

	if tx.Nonce == "" {
		reqGetNonce := &types.GetAddrNonceByHashArgs{
			Address: tx.From,
		}
		var nonce *int64
		if err := conn.cli.CallMethod(1, "TxPool.GetAddrTxNonce", &reqGetNonce, &nonce); err != nil {
			return "", fmt.Errorf("invalid GetAddrTxNonce addr:%v err:%v", tx.From, err)
		}
		tx.Nonce = strconv.FormatInt(*nonce, 10)
	}

	if tx.GasLimit == "" {
		tx.GasLimit = common.TxGas.String()
	} else {
		gaslimit, ok := new(big.Int).SetString(tx.GasLimit, 10)
		if !ok {
			return "", fmt.Errorf("failed to parse gaslimit")
		}
		if gaslimit.Cmp(common.TxGas) < 0 {
			return "", fmt.Errorf("gaslimit did not reach the lowest peugeot")
		}
	}

	if tx.GasLimit == "" {
		tx.GasLimit = common.TxGas.String()
	} else {
		gaslimit, ok := new(big.Int).SetString(tx.GasLimit, 10)
		if !ok {
			return "", fmt.Errorf("failed to parse gaslimit")
		}
		if gaslimit.Cmp(common.TxGas) < 0 {
			return "", fmt.Errorf("gaslimit did not reach the lowest peugeot")
		}
		tx.GasLimit = gaslimit.Text(10)
	}

	if tx.GasPrice == "" {
		tx.GasPrice = common.DefaultGasPrice().String()
	} else {
		gasprice, ok := new(big.Int).SetString(tx.GasPrice, 10)
		if !ok {
			return "", fmt.Errorf("failed to parse gaslimit")
		}
		if gasprice.Cmp(common.DefaultGasPrice()) < 0 {
			return "", fmt.Errorf("gasprice did not reach the lowest peugeot")
		}
		tx.GasLimit = gasprice.Text(10)
	}
	if err := tx.SignWithPrivateKey("0x010152997afc9ae613717a58fb5a8c21617de5ad73576e6afaba7c8b3f4319b08ca0"); err != nil {
		return "", err
	}

	result, err := tx.Transfer2Raw()
	if err != nil {
		return "", err
	}

	return result, nil
}

func (conn *Connection) SendRawTransaction(rawTx string) (string, error) {
	tx := types.RawTransactionArgs{
		Data: rawTx,
	}

	var result string
	err := conn.cli.CallMethod(1, "TxPool.SendRawTransaction", &tx, &result)
	if err != nil {

		return "", nil
	}

	return result, nil
}

func (conn *Connection) checkChainId(expected uint8) error {
	return nil
}

func (conn *Connection) Client() *rpcclient.Client {
	return conn.cli
}
