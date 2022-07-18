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
	// var prop *proposal
	// var err error

	// // Construct the proposal
	// switch m.Type {
	// case msg.FungibleTransfer:
	// 	prop, err = w.createFungibleProposal(m)
	// case msg.NonFungibleTransfer:
	// 	prop, err = w.createNonFungibleProposal(m)
	// case msg.GenericTransfer:
	// 	prop, err = w.createGenericProposal(m)
	// default:
	// 	w.sysErr <- fmt.Errorf("unrecognized message type received (chain=%d, name=%s)", m.Destination, w.conn.name)
	// 	return false
	// }

	// if err != nil {
	// 	w.sysErr <- fmt.Errorf("failed to construct proposal (chain=%d, name=%s) Error: %w", m.Destination, w.conn.name, err)
	// 	return false
	// }

	// for i := 0; i < BlockRetryLimit; i++ {
	// 	// Ensure we only submit a vote if the proposal hasn't completed
	// 	valid, reason, err := w.proposalValid(prop)
	// 	if err != nil {
	// 		w.log.Error("Failed to assert proposal state", "err", err)
	// 		time.Sleep(BlockRetryInterval)
	// 		continue
	// 	}

	// 	// If active submit call, otherwise skip it. Retry on failure.
	// 	if valid {
	// 		w.log.Info("Acknowledging proposal on chain", "nonce", prop.depositNonce, "source", prop.sourceId, "resource", fmt.Sprintf("%x", prop.resourceId), "method", prop.method)

	// 		err = w.conn.SubmitTx(AcknowledgeProposal, prop.depositNonce, prop.sourceId, prop.resourceId, prop.call)
	// 		if err != nil && err.Error() == TerminatedError.Error() {
	// 			return false
	// 		} else if err != nil {
	// 			w.log.Error("Failed to execute extrinsic", "err", err)
	// 			time.Sleep(BlockRetryInterval)
	// 			continue
	// 		}
	// 		if w.metrics != nil {
	// 			w.metrics.VotesSubmitted.Inc()
	// 		}
	// 		return true
	// 	} else {
	// 		w.log.Info("Ignoring proposal", "reason", reason, "nonce", prop.depositNonce, "source", prop.sourceId, "resource", prop.resourceId)
	// 		return true
	// 	}
	// }
	return true
}
