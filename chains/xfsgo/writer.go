package xfsgo

import (
	"errors"

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

	w.log.Info("ResolveErc20", "m.Payload", m.Payload)

	w.conn.TransferIn("", "", "", 0, 0, "", 10)

	return true
}
