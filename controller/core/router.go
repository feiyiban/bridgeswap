package core

import (
	"bridgeswap/controller/msg"
	log "bridgeswap/logger"
	"fmt"
	"sync"
)

// Writer consumes a message and makes the requried on-chain interactions.
type Writer interface {
	ResolveMessage(message msg.Message) bool
}

// Router forwards messages from their source to their destination
type Router struct {
	registry map[uint8]Writer
	lock     *sync.RWMutex
	log      log.Logger
}

func NewRouter(log log.Logger) *Router {
	return &Router{
		registry: make(map[uint8]Writer),
		lock:     &sync.RWMutex{},
		log:      log,
	}
}

// Listen registers a Writer with a ChainId which Router.Send can then use to propagate messages
func (r *Router) Listen(id uint8, w Writer) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.log.Debug("Registering new chain in router", "ChainId", id)
	r.registry[id] = w
}

// Send passes a message to the destination Writer if it exists
func (r *Router) Send(msg msg.Message) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.log.Trace("Routing message", "src", msg.Source, "dest", msg.Destination)
	w := r.registry[msg.Destination]
	if w == nil {
		return fmt.Errorf("unknown destination chainId: %d", msg.Destination)
	}

	go w.ResolveMessage(msg)
	return nil
}
