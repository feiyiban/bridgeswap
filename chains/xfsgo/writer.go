package xfsgo

import (
	"bridgeswap/bindings/xfs/xfsbridge"
	"bridgeswap/chains/xfsgo/types"
	"bridgeswap/controller/msg"
	"bridgeswap/logger"
	"bridgeswap/sdk/xfsgo/common"
	"fmt"
	"math/big"
)

type writer struct {
	cfg    Config
	conn   Connection
	sysErr chan<- error
	log    logger.Logger
}

func NewWriter(cfg *Config, conn Connection, sysErr chan<- error, log logger.Logger) *writer {
	return &writer{
		cfg:    *cfg,
		conn:   conn,
		sysErr: sysErr,
		log:    log,
	}
}

func (write *writer) start() error {
	write.log.Info("Starting xfsgo writer...")
	return nil
}

func (write *writer) ResolveMessage(message msg.Message) bool {
	write.log.Info("Attempting to resolve message", "type", message.Type, "src", message.Source, "dst", message.Destination)

	switch message.Type {
	case msg.TokenTransfer:
		return write.ResolveERCToken(message)
	default:
		write.log.Error("Unknown message type received", "type", message.Type)
		return false
	}
}

func (write *writer) ResolveERCToken(message msg.Message) bool {
	write.log.Info("XFSGO", "ResolveERCToken msg.Payload", message.Payload)

	if len(message.Payload) <= 0 {
		write.log.Error("XFSGO", "ResolveERCToken message payload", "is nil")
		return false
	}

	write.log.Debug("XFSGO", "ResolveERCToken message payload", message.Payload)

	fromChainId := big.NewInt(0).SetUint64(uint64(message.Source))
	strDestAddr := message.Payload[0].(string)
	strValue := message.Payload[1].(string)
	tokenAddr := write.cfg.erc20Contract
	fromAddr := write.cfg.from
	write.log.Info("TransferIn asset to XFSGO chain", "tokenAddr", tokenAddr, "fromAddr", fromAddr, "strDestAddr", strDestAddr, "value", strValue)
	amount, ok := new(big.Int).SetString(strValue, 10)
	if !ok {
		write.log.Error("XFSGO", "ResolveERCToken TransferIn", "Abnormal asset value")
		return false
	}

	tokenABI, err := xfsbridge.JSON(xfsbridge.BRIDGETOKENABI)
	if err != nil {
		write.log.Error("XFSGO", "ResolveERCToken TransferIn", fmt.Errorf("%v parse json err", xfsbridge.BRIDGETOKENABI))
		return false
	}

	cAddr := common.StrB58ToAddress(strDestAddr)
	TransferInAbiData, err := tokenABI.TransferIn(xfsbridge.NewAddress(cAddr), xfsbridge.NewUint256(amount), xfsbridge.NewUint256(fromChainId))
	if err != nil {
		write.log.Error("XFSGO", "ResolveERCToken TransferIn", err)
		return false
	}

	rawTx := types.StringRawTransaction{
		From: write.cfg.from,
		To:   write.cfg.bridgeContract,
		Data: TransferInAbiData,
	}
	strRawTx, err := write.conn.SignedTx(rawTx)
	if err != nil {
		write.log.Error("XFSGO", "ResolveERCToken TransferIn SignedTx", err)
		return false
	}

	txHash, err := write.conn.SendRawTransaction(strRawTx)
	if err != nil {
		write.log.Error("XFSGO", "ResolveERCToken TransferIn SendRawTransaction", err)
		return false
	}

	write.log.Info("XFSGO ", "ResolveERCToken successfully", txHash)

	return true
}
