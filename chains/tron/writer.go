package tron

import (
	"bridgeswap/controller/core"
	"bridgeswap/controller/msg"
	"bytes"
	"encoding/json"
	"math/big"

	"bridgeswap/logger"

	tronabi "bridgeswap/chains/tron/pkg/abi"
)

var _ core.Writer = &writer{}

var PassedStatus uint8 = 2
var TransferredStatus uint8 = 3
var CancelledStatus uint8 = 4

type writer struct {
	cfg            Config
	conn           Connection
	bridgeContract string // instance of bound receiver bridgeContract
	log            logger.Logger
	stop           <-chan int
	sysErr         chan<- error // Reports fatal error to core
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

func (w *writer) start() error {
	w.log.Debug("Starting tron writer...")
	return nil
}

// setContract adds the bound receiver bridgeContract to the writer
func (w *writer) setContract(bridge string) {
	w.bridgeContract = bridge
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
	byteValue := new(bytes.Buffer)
	byteValue.Write(m.Payload[32:64])
	toAddr := new(bytes.Buffer)
	toAddr.Write(m.Payload[96:130])
	w.log.Info("Creating erc20", "src", m.Source, "byteValue", byteValue.Bytes(), "toaddr", toAddr.Bytes())
	addr := toAddr.Bytes()

	tokenAddr := w.cfg.erc20Contract.String()
	fromAddr := w.cfg.from
	destAddr := string(addr)
	w.log.Info("Depositout Tron", "tokenAddr", tokenAddr, "fromAddr", fromAddr, "destAddr", destAddr, "value", big.NewInt(0).SetBytes(byteValue.Bytes()))

	param := []tronabi.Param{
		{"address": tokenAddr},
		{"address": destAddr},
		{"uint256": big.NewInt(0).SetBytes(byteValue.Bytes()).String()},
		{"uint256": big.NewInt(int64(m.Source)).String()},
		{"uint256": big.NewInt(int64(m.Destination)).String()},
	}

	w.log.Info("hex.EncodeToString", "value", big.NewInt(0).SetBytes(byteValue.Bytes()).String())
	dataBuf, err := json.Marshal(param)
	if err != nil {
		return false
	}
	err = w.conn.TransferIn(w.cfg.from, w.bridgeContract, string(dataBuf), int64(4000000000), 0, "", 0)
	if err != nil {
		return false
	}

	w.log.Info("Tron Bridge", "ExecuteTransaction", "successfully")
	return true
}
