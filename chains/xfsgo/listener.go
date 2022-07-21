package xfsgo

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"bridgeswap/bindings/xfs/events"
	"bridgeswap/bindings/xfs/xfsbridge"
	"bridgeswap/blockstore"
	"bridgeswap/chains"
	"bridgeswap/chains/xfsgo/types"
	"bridgeswap/controller/msg"

	"bridgeswap/logger"
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
	router     chains.Router
	log        logger.Logger
	stop       <-chan int
	sysErr     chan<- error
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

	// querying for logs
	logRequst := types.GetLogsRequest{
		FromBlock: latestBlock.String(),
		ToBlock:   latestBlock.String(),
		Address:   l.cfg.bridgeContract,
		EventHash: xfsbridge.TransoutEvent,
	}
	logs, err := l.conn.GetLogs(logRequst)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %w", err)
	}

	for _, log := range *logs {
		l.log.Info("getDepositEventsForBlock", "result", log)

		if log.EventHash.Hex() == xfsbridge.TransoutEvent {
			event, err := events.JSON(xfsbridge.BRIDGETOKENABI)
			if err != nil {
				return err
			}

			val, err := event.Decode(xfsbridge.TransoutEvent, strings.TrimPrefix(log.EventValue, "0x"))
			if err != nil {
				return err
			}

			fromChainID := val["fromChainId"].(string)

			fromID, ok := new(big.Int).SetString(fromChainID, 10)
			if !ok {
				return err
			}
			toChainID := val["toChainId"].(string)
			toID, ok := new(big.Int).SetString(toChainID, 10)
			if !ok {
				return err
			}
			toAddr := val["toAddress"].(string)
			value := val["value"].(string)

			m := msg.Message{
				Source:      uint8(fromID.Uint64()),
				Destination: uint8(toID.Uint64()),
				Type:        msg.TokenTransfer,
				Payload: []interface{}{
					toAddr,
					value,
				},
			}

			if err != nil {
				return err
			}

			l.log.Info("message", "source:", m.Source, "destination:", m.Destination, "toAddr:", toAddr, "value", value)
			err = l.router.Send(m)
			if err != nil {
				l.log.Error("subscription error: failed to route message", "err", err)
			}

		}
	}

	return nil
}
