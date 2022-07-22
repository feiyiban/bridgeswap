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
	sysErr chan<- error       // Reports fatal error to core
	stop   <-chan int
	log    logger.Logger
}

// NewWriter creates and returns writer
func NewWriter(cfg *Config, conn Connection, sysErr chan<- error, stop <-chan int, log logger.Logger) *writer {
	return &writer{
		cfg:    *cfg,
		conn:   conn,
		sysErr: sysErr,
		stop:   stop,
		log:    log,
	}
}

// setContract adds the bound receiver bridgeContract to the writer
func (write *writer) setContract(bridge *bridgev1.Bridgev1) {
	write.bridge = bridge
}

func (write *writer) start() error {
	write.log.Debug("Starting ethereum writer...")
	return nil
}

// ResolveMessage handles any given message based on type
// A bool is returned to indicate failure/success, this should be ignored except for within tests.
func (write *writer) ResolveMessage(m msg.Message) bool {
	write.log.Info("Attempting to resolve message", "type", m.Type, "src", m.Source, "dst", m.Destination)

	switch m.Type {
	case msg.TokenTransfer:
		return write.ResolveERCToken(m)
	default:
		write.log.Error("Unknown message type received", "type", m.Type)
		return false
	}
}

func (write *writer) ResolveERCToken(m msg.Message) bool {

	write.log.Info("ResolveErc20", "m.Payload", m.Payload)

	if len(m.Payload) <= 0 {
		return false
	}

	toAddr := m.Payload[0].(string)
	value := m.Payload[1].(string)

	tokenAddr := write.cfg.erc20Contract
	fromAddr := write.cfg.from
	write.log.Info("Depositout Eth", "tokenAddr", tokenAddr, "fromAddr", fromAddr, "destAddr", toAddr, "value", value)

	amout, _ := big.NewInt(0).SetString(value, 10)

	nonce, err := write.conn.Client().NonceAt(context.Background(), common.HexToAddress(fromAddr), nil)
	if err != nil {

		return false
	}
	gasPrice, err := write.conn.Client().SuggestGasPrice(context.Background())
	if err != nil {
		return false
	}

	signer := types.HomesteadSigner{}
	auth := &bind.TransactOpts{
		From:     common.HexToAddress(fromAddr),
		Nonce:    new(big.Int).SetUint64(nonce),
		GasPrice: gasPrice,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			keyPair := write.conn.Keypair()
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

	err = write.BridgeTransferIn(auth, common.HexToAddress(write.cfg.erc20Contract), common.HexToAddress(toAddr), amout, big.NewInt(0).SetUint64(uint64(m.Source)), big.NewInt(0).SetUint64(uint64(m.Destination)))

	return err == nil
}
