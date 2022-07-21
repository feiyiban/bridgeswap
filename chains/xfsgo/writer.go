package xfsgo

import (
	"errors"
	"math/big"

	"bridgeswap/bindings/xfs/xfsbridge"
	"bridgeswap/chains/xfsgo/types"
	"bridgeswap/controller/core"
	"bridgeswap/sdk/xfs/common"

	"bridgeswap/controller/msg"
	"bridgeswap/logger"
)

var _ core.Writer = &writer{}

var TerminatedError = errors.New("terminated")

type writer struct {
	cfg        Config
	conn       Connection
	log        logger.Logger
	sysErr     chan<- error
	extendCall bool // Extend extrinsic calls to substrate with ResourceID.Used for backward compatibility with example pallet.
}

func NewWriter(conn Connection, cfg *Config, log logger.Logger, sysErr chan<- error, extendCall bool) *writer {
	return &writer{
		cfg:        *cfg,
		conn:       conn,
		log:        log,
		sysErr:     sysErr,
		extendCall: extendCall,
	}
}

func (w *writer) start() error {
	w.log.Debug("Starting tron writer...")
	return nil
}

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

	fromId := big.NewInt(0).SetUint64(uint64(m.Source))
	toAddr := m.Payload[0].(string)
	value := m.Payload[1].(string)

	tokenAddr := w.cfg.erc20Contract
	fromAddr := w.cfg.from
	destAddr := string(toAddr)
	w.log.Info("Depositout XFS", "tokenAddr", tokenAddr, "fromAddr", fromAddr, "destAddr", destAddr, "value", value)

	w.log.Info("ResolveErc20", "m.Payload", m.Payload)

	amount, ok := new(big.Int).SetString(value, 10)
	if !ok {
		return false
	}

	abi, err := xfsbridge.JSON(xfsbridge.BRIDGETOKENABI)
	if err != nil {
		return false
	}
	cAddr := common.StrB58ToAddress(toAddr)
	data, err := abi.TransferIn(xfsbridge.NewAddress(cAddr), xfsbridge.NewUint256(amount), xfsbridge.NewUint256(fromId))
	if err != nil {
		return false
	}

	rawTx := types.StringRawTransaction{
		From: w.cfg.from,
		To:   w.cfg.bridgeContract,
		Data: data,
	}
	strRawTx, err := w.conn.SignedTx(rawTx)
	if err != nil {
		return false
	}

	txHash, err := w.conn.SendRawTransaction(strRawTx)

	if err != nil {
		return false
	}

	w.log.Info("XFS", "SendRawTransaction successfully", txHash)

	return true
}
