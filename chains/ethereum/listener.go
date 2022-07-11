package ethereum

import (
	"bridgeswap/blockstore"
	"bridgeswap/logger"
	"errors"
	"math/big"
	"time"

	"bridgeswap/bindings/eth/bridgev1"
	"bridgeswap/chains"
)

var (
	BlockRetryInterval = time.Second * 5
	BlockRetryLimit    = 5
	ErrFatalPolling    = errors.New("listener block polling failed")
)

type listener struct {
	cfg        Config
	conn       Connection
	router     chains.Router
	bridge     *bridgev1.Bridgev1 // instance of bound bridge contract
	blockstore blockstore.Blockstorer
	log        logger.Logger
	stop       <-chan int
	sysErr     chan<- error // Reports fatal error to core

	blockConfirmations *big.Int
}

func NewListener(conn Connection, cfg *Config, log logger.Logger, bs blockstore.Blockstorer, stop <-chan int, sysErr chan<- error) *listener {
	return &listener{
		cfg:                *cfg,
		conn:               conn,
		log:                log,
		blockstore:         bs,
		stop:               stop,
		sysErr:             sysErr,
		blockConfirmations: cfg.blockConfirmations,
	}
}

// setContracts sets the listener with the appropriate contracts
func (listen *listener) setContracts(bridge *bridgev1.Bridgev1) {
	listen.bridge = bridge
}

// sets the router
func (listen *listener) setRouter(router chains.Router) {
	listen.router = router
}

// start registers all subscriptions provided by the config
func (listen *listener) start() error {
	listen.log.Debug("Starting listener...")

	go func() {
		err := listen.pollBlocks()
		if err != nil {
			listen.log.Error("Polling blocks failed", "err", err)
		}
	}()

	return nil
}

// pollBlocks will poll for the latest block and proceed to parse the associated events as it sees new blocks.
// Polling begins at the block defined in `l.cfg.startBlock`. Failed attempts to fetch the latest block or parse
// a block will be retried up to BlockRetryLimit times before continuing to the next block.
func (listen *listener) pollBlocks() error {
	var currentBlock = listen.cfg.startBlock
	listen.log.Info("Polling Blocks...", "block", currentBlock)

	var retry = BlockRetryLimit
	for {
		select {
		case <-listen.stop:
			return errors.New("polling terminated")
		default:
			// No more retries, goto next block
			if retry == 0 {
				listen.log.Error("Polling failed, retries exceeded")
				listen.sysErr <- ErrFatalPolling
				return nil
			}

			latestBlock, err := listen.conn.LatestBlock()
			if err != nil {
				listen.log.Error("Unable to get latest block", "block", currentBlock, "err", err)
				retry--
				time.Sleep(BlockRetryInterval)
				continue
			}

			// if listen.metrics != nil {
			// 	listen.metrics.LatestKnownBlock.Set(float64(latestBlock.Int64()))
			// }

			// Sleep if the difference is less than BlockDelay; (latest - current) < BlockDelay
			if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(listen.blockConfirmations) == -1 {
				listen.log.Debug("Block not ready, will retry", "target", currentBlock, "latest", latestBlock)
				time.Sleep(BlockRetryInterval)
				continue
			}

			// Parse out events
			err = listen.getDepositEventsForBlock(currentBlock)
			if err != nil {
				listen.log.Error("Failed to get events for block", "block", currentBlock, "err", err)
				retry--
				continue
			}

			// Write to block store. Not a critical operation, no need to retry
			err = listen.blockstore.StoreBlock(currentBlock)
			if err != nil {
				listen.log.Error("Failed to write latest block to blockstore", "block", currentBlock, "err", err)
			}

			// if listen.metrics != nil {
			// 	listen.metrics.BlocksProcessed.Inc()
			// 	listen.metrics.LatestProcessedBlock.Set(float64(latestBlock.Int64()))
			// }

			// listen.latestBlock.Height = big.NewInt(0).Set(latestBlock)
			// listen.latestBlock.LastUpdated = time.Now()

			// Goto next block and reset retry counter
			currentBlock.Add(currentBlock, big.NewInt(1))
			retry = BlockRetryLimit
		}
	}
}

// getDepositEventsForBlock looks for the deposit event in the latest block
func (l *listener) getDepositEventsForBlock(latestBlock *big.Int) error {
	l.log.Debug("Querying block for deposit events", "block", latestBlock)

	return nil
}
