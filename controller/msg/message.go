package msg

var (
	TokenTransfer string = "TokenTransfer"
)

type Message struct {
	Source      uint8  // Source where message was initiated
	Destination uint8  // Destination chain of message
	Type        string // type of bridge transfer
	Payload     []byte // data associated with event sequence
}
