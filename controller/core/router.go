package core

import (
	"bridgeswap/controller/msg"
	log "bridgeswap/logger"
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
