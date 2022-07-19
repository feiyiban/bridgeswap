package xfsgo

import (
	"errors"
	"math/big"

	"bridgeswap/bindings/xfs/xfsbridge"
	"bridgeswap/chains/xfsgo/pkg/common"
	"bridgeswap/chains/xfsgo/pkg/types"
	"bridgeswap/controller/core"

	"bridgeswap/controller/msg"
	"bridgeswap/logger"
)

var _ core.Writer = &writer{}

var TerminatedError = errors.New("terminated")

type writer struct {
	conn       Connection
	log        logger.Logger
	sysErr     chan<- error
	extendCall bool // Extend extrinsic calls to substrate with ResourceID.Used for backward compatibility with example pallet.
}

func NewWriter(conn Connection, log logger.Logger, sysErr chan<- error, extendCall bool) *writer {
	return &writer{
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

	var depositorAddress = ""
	var amount = ""
	var fromChainID = ""

	w.log.Info("ResolveErc20", "m.Payload", m.Payload)

	value, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return false
	}

	chainID, ok := new(big.Int).SetString(fromChainID, 10)
	if !ok {
		return false
	}

	abi, err := xfsbridge.JSON(xfsbridge.BRIDGETOKENABI)
	addr := common.StrB58ToAddress(depositorAddress)
	cTypeAddr := xfsbridge.NewAddress(addr)

	data, err := abi.TransferIn(cTypeAddr, xfsbridge.NewUint256(value), xfsbridge.NewUint256(chainID))
	if err != nil {
		return false
	}

	rawTx := types.StringRawTransaction{
		Data: data,
	}
	strRawTx, err := w.conn.SignedTx(rawTx)
	if err != nil {
		return false
	}

	w.conn.SendRawTransaction(*strRawTx)

	w.conn.TransferIn("", "", "", 0, 0, "", 10)

	return true
}
