package ethereum

import (
	"bridgeswap/bindings/eth/bridgev1"
	"bridgeswap/controller/msg"
	"bridgeswap/logger"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	TokenTransfer string = "TokenTransfer"
)

type writer struct {
	cfg    Config
	conn   Connection
	bridge *bridgev1.Bridgev1 // instance of bound receiver bridgeContract
	log    logger.Logger
	stop   <-chan int
	sysErr chan<- error // Reports fatal error to core
}

// NewWriter creates and returns writer
func NewWriter(conn Connection, cfg *Config, log logger.Logger, stop <-chan int, sysErr chan<- error) *writer {
	return &writer{
		cfg:    *cfg,
		conn:   conn,
		log:    log,
		stop:   stop,
		sysErr: sysErr,
	}
}

// setContract adds the bound receiver bridgeContract to the writer
func (w *writer) setContract(bridge *bridgev1.Bridgev1) {
	w.bridge = bridge
}

func (write *writer) start() error {
	write.log.Debug("Starting ethereum writer...")
	return nil
}

// ResolveMessage handles any given message based on type
// A bool is returned to indicate failure/success, this should be ignored except for within tests.
func (w *writer) ResolveMessage(m msg.Message) bool {
	w.log.Info("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination)

	switch m.Type {
	case msg.TokenTransfer:
		return w.ResolveErc20(m)
	default:
		w.log.Error("Unknown message type received", "type", m.Type)
		return false
	}
}

func (w *writer) ResolveErc20(m msg.Message) bool {

	w.log.Info("ResolveErc20", "m.Payload", m.Payload)

	if len(m.Payload) <= 0 {
		return false
	}

	toAddr := m.Payload[0].(string)
	value := m.Payload[1].(string)

	tokenAddr := w.cfg.erc20Contract.String()
	fromAddr := w.cfg.from
	w.log.Info("Depositout Eth", "tokenAddr", tokenAddr, "fromAddr", fromAddr, "destAddr", toAddr, "value", value)

	amout, _ := big.NewInt(0).SetString(value, 10)

	nonce, err := w.conn.Client().NonceAt(context.Background(), common.HexToAddress(fromAddr), nil)
	if err != nil {

		return false
	}
	gasPrice, err := w.conn.Client().SuggestGasPrice(context.Background())
	if err != nil {
		return false
	}

	signer := types.HomesteadSigner{}
	auth := &bind.TransactOpts{
		From:     common.HexToAddress(fromAddr),
		Nonce:    new(big.Int).SetUint64(nonce),
		GasPrice: gasPrice,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			keyPair := w.conn.Keypair()
			if keyPair == nil {
				return nil, fmt.Errorf("%s can't find", fromAddr)
			}
			signature, _ := crypto.Sign(signer.Hash(tx).Bytes(), keyPair.PrivateKey())
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}

	err = w.BridgeTransferIn(auth, w.cfg.erc20Contract, common.HexToAddress(toAddr), amout, big.NewInt(0).SetUint64(uint64(m.Source)), big.NewInt(0).SetUint64(uint64(m.Destination)))

	return err == nil
}
