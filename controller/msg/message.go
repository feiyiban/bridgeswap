package msg

import "math/big"

var (
	TokenTransfer string = "TokenTransfer"
)

type Message struct {
	Source      uint8         // Source where message was initiated
	Destination uint8         // Destination chain of message
	Type        string        // type of bridge transfer
	Payload     []interface{} // data associated with event sequence
}

func NewErc20TransferOut(source, dest uint8, amount *big.Int, toAddr string) Message {
	return Message{
		Source:      source,
		Destination: dest,
		Type:        TokenTransfer,
		Payload: []interface{}{
			amount.Bytes(),
			toAddr,
		},
	}
}
