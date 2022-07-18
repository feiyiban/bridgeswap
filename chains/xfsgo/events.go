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
	// evt, ok := evtI.(events.EventGenericTransfer)
	// if !ok {
	// 	return msg.Message{}, fmt.Errorf("failed to cast EventGenericTransfer type")
	// }

	// log.Info("Got generic transfer event!", "destination", evt.Destination, "resourceId", evt.ResourceId)

	// return msg.NewGenericTransfer(
	// 	0, // Unset
	// 	msg.ChainId(evt.Destination),

	// ), nil
	return msg.Message{}, nil
}
