package xfsgo

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"bridgeswap/blockstore"
	"bridgeswap/chains"

	"bridgeswap/logger"

	"bridgeswap/chains/xfsgo/pkg/types"
)

var (
	// Frequency of polling for a new block
	BlockRetryInterval = time.Second * 5
	BlockRetryLimit    = 5
)

type listener struct {
	cfg        Config
	conn       Connection
	blockstore blockstore.Blockstorer

	router chains.Router
	log    logger.Logger
	stop   <-chan int
	sysErr chan<- error
}

func NewListener(conn Connection, cfg *Config, log logger.Logger, bs blockstore.Blockstorer, stop <-chan int, sysErr chan<- error) *listener {
	return &listener{
		cfg:        *cfg,
		blockstore: bs,
		conn:       conn,
		log:        log,
		stop:       stop,
		sysErr:     sysErr,
	}
}

func (l *listener) setRouter(r chains.Router) {
	l.router = r
}

// start creates the initial subscription for all events
func (l *listener) start() error {
	go func() {
		err := l.pollBlocks()
		if err != nil {
			l.log.Error("Polling blocks failed", "err", err)
		}
	}()

	return nil
}

var ErrBlockNotReady = errors.New("required result to be 32 bytes, but got 0")

// pollBlocks will poll for the latest block and proceed to parse the associated events as it sees new blocks.
// Polling begins at the block defined in `l.startBlock`. Failed attempts to fetch the latest block or parse
// a block will be retried up to BlockRetryLimit times before returning with an error.
func (listen *listener) pollBlocks() error {
	var currentBlock = listen.cfg.startBlock
	var retry = BlockRetryLimit
	for {
		select {
		case <-listen.stop:
			return errors.New("terminated")
		default:
			// No more retries, goto next block
			if retry == 0 {
				listen.sysErr <- fmt.Errorf("event polling retries exceeded (chain=%d, name=%s)", listen.cfg.id, listen.cfg.name)
				return nil
			}

			latestBlock, err := listen.conn.LatestBlock()
			if err != nil {
				listen.log.Error("Unable to get latest block", "block", currentBlock, "err", err)
				retry--
				time.Sleep(BlockRetryInterval)
				continue
			}

			if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(listen.cfg.blockConfirmations) == -1 {
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

			err = listen.blockstore.StoreBlock(currentBlock)
			if err != nil {
				listen.log.Error("Failed to write latest block to blockstore", "block", currentBlock, "err", err)
			}

			// Goto next block and reset retry counter
			currentBlock.Add(currentBlock, big.NewInt(1))
			retry = BlockRetryLimit

		}
	}
}

// getDepositEventsForBlock looks for the deposit event in the latest block
func (l *listener) getDepositEventsForBlock(latestBlock *big.Int) error {
	l.log.Debug("Querying block for deposit events", "block", latestBlock)
	// query := buildQuery(l.cfg.bridgeContract, utils.Deposit, latestBlock, latestBlock)

	// // querying for logs
	logRequst := types.GetLogsRequest{}
	logs, err := l.conn.GetLogs(logRequst)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %w", err)
	}

	l.log.Info("getDepositEventsForBlock", "result", logs)

	// // read through the log events and handle their deposit event if handler is recognized
	// for _, log := range logs {
	// 	var m msg.Message
	// 	destId := msg.ChainId(log.Topics[1].Big().Uint64())
	// 	rId := msg.ResourceIdFromSlice(log.Topics[2].Bytes())
	// 	nonce := msg.Nonce(log.Topics[3].Big().Uint64())

	// 	addr, err := l.bridgeContract.ResourceIDToHandlerAddress(&bind.CallOpts{From: l.conn.Keypair().CommonAddress()}, rId)
	// 	if err != nil {
	// 		return fmt.Errorf("failed to get handler from resource ID %x", rId)
	// 	}

	// 	if addr == l.cfg.erc20HandlerContract {
	// 		m, err = l.handleErc20DepositedEvent(destId, nonce)
	// 	} else if addr == l.cfg.erc721HandlerContract {
	// 		m, err = l.handleErc721DepositedEvent(destId, nonce)
	// 	} else if addr == l.cfg.genericHandlerContract {
	// 		m, err = l.handleGenericDepositedEvent(destId, nonce)
	// 	} else {
	// 		l.log.Error("event has unrecognized handler", "handler", addr.Hex())
	// 		return nil
	// 	}

	// 	if err != nil {
	// 		return err
	// 	}

	// 	err = l.router.Send(m)
	// 	if err != nil {
	// 		l.log.Error("subscription error: failed to route message", "err", err)
	// 	}
	// }

	return nil
}
