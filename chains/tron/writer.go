package tron

import (
	"bridgeswap/controller/core"
	"bridgeswap/controller/msg"
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

func (write *writer) start() error {
	write.log.Debug("Starting tron writer...")
	return nil
}

// setContract adds the bound receiver bridgeContract to the writer
func (write *writer) setContract(bridge string) {
	write.bridgeContract = bridge
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

	tokenAddr := write.cfg.erc20Contract.String()
	fromAddr := write.cfg.from
	write.log.Info("Depositout Tron", "tokenAddr", tokenAddr, "fromAddr", fromAddr, "destAddr", toAddr, "value", value)

	param := []tronabi.Param{
		{"address": tokenAddr},
		{"address": toAddr},
		{"uint256": value},
		{"uint256": big.NewInt(int64(m.Source)).String()},
		{"uint256": big.NewInt(int64(m.Destination)).String()},
	}

	write.log.Info("hex.EncodeToString", "value", value)
	dataBuf, err := json.Marshal(param)
	if err != nil {
		return false
	}
	err = write.conn.TransferIn(write.cfg.from, write.bridgeContract, string(dataBuf), int64(4000000000), 0, "", 0)
	if err != nil {
		return false
	}

	write.log.Info("Tron Bridge", "ExecuteTransaction", "successfully")
	return true
}
