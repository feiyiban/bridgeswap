package chains

import "bridgeswap/controller/msg"

type Router interface {
	Send(message msg.Message) error
}
