package xfsgo

import (
	"bridgeswap/controller/msg"

	"bridgeswap/logger"
)

type eventName string
type eventHandler func(interface{}, logger.Logger) (msg.Message, error)

const FungibleTransfer eventName = "FungibleTransfer"
const NonFungibleTransfer eventName = "NonFungibleTransfer"
const GenericTransfer eventName = "GenericTransfer"

var Subscriptions = []struct {
	name    eventName
	handler eventHandler
}{
	{GenericTransfer, genericTransferHandler},
}

func genericTransferHandler(evtI interface{}, log logger.Logger) (msg.Message, error) {

	return msg.Message{}, nil
}
