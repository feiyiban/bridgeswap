package ethereum

import (
	"bridgeswap/bindings/eth/bridgev1"
	"bridgeswap/controller/msg"
	"bridgeswap/logger"
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
		// return w.createGenericDepositProposal(m)
		return true
	default:
		w.log.Error("Unknown message type received", "type", m.Type)
		return false
	}
}
