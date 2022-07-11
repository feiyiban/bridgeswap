package tron

import (
	"errors"
	"math/big"
	"time"

	"bridgeswap/blockstore"
	// metrics "crosschainbridge/controller/metrics/types"

	"bridgeswap/chains"

	"bridgeswap/logger"

	eth "github.com/ethereum/go-ethereum"
	ethcommon "github.com/ethereum/go-ethereum/common"
)

var BlockRetryInterval = time.Second * 5
var BlockRetryLimit = 5
var ErrFatalPolling = errors.New("listener block polling failed")

type listener struct {
	cfg    Config
	conn   Connection
	router chains.Router
	// bridgeContract     *bridgemap.Bridge // instance of bound bridge contract
	log        logger.Logger
	blockstore blockstore.Blockstorer
	stop       <-chan int
	sysErr     chan<- error // Reports fatal error to core
	// latestBlock        metrics.LatestBlock
	// metrics            *metrics.ChainMetrics
	blockConfirmations *big.Int
}

// NewListener creates and returns a listener
func NewListener(conn Connection, cfg *Config, log logger.Logger, stop <-chan int, sysErr chan<- error) *listener {
	return &listener{
		cfg:    *cfg,
		conn:   conn,
		log:    log,
		stop:   stop,
		sysErr: sysErr,
		// latestBlock:        metrics.LatestBlock{LastUpdated: time.Now()},
		blockConfirmations: cfg.blockConfirmations,
	}
}

// setContracts sets the listener with the appropriate contracts
// func (l *listener) setContracts(bridge *bridgemap.Bridge) {
// 	l.bridgeContract = bridge
// }

// sets the router
func (l *listener) setRouter(r chains.Router) {
	l.router = r
}

// start registers all subscriptions provided by the config
func (l *listener) start() error {
	l.log.Debug("Starting listener...")

	// go func() {
	// 	err := l.pollBlocks()
	// 	if err != nil {
	// 		l.log.Error("Polling blocks failed", "err", err)
	// 	}
	// }()

	return nil
}

// pollBlocks will poll for the latest block and proceed to parse the associated events as it sees new blocks.
// Polling begins at the block defined in `l.cfg.startBlock`. Failed attempts to fetch the latest block or parse
// a block will be retried up to BlockRetryLimit times before continuing to the next block.
func (l *listener) pollBlocks() error {
	var currentBlock = l.cfg.startBlock
	l.log.Info("Polling Blocks...", "block", currentBlock)

	var retry = BlockRetryLimit
	for {
		select {
		case <-l.stop:
			return errors.New("polling terminated")
		default:
			// No more retries, goto next block
			if retry == 0 {
				l.log.Error("Polling failed, retries exceeded")
				l.sysErr <- ErrFatalPolling
				return nil
			}

			latestBlock, err := l.conn.LatestBlock()
			if err != nil {
				l.log.Error("Unable to get latest block", "block", currentBlock, "err", err)
				retry--
				time.Sleep(BlockRetryInterval)
				continue
			}

			// if l.metrics != nil {
			// 	l.metrics.LatestKnownBlock.Set(float64(latestBlock.Int64()))
			// }

			// Sleep if the difference is less than BlockDelay; (latest - current) < BlockDelay
			if big.NewInt(0).Sub(latestBlock, currentBlock).Cmp(l.blockConfirmations) == -1 {
				l.log.Debug("Block not ready, will retry", "target", currentBlock, "latest", latestBlock)
				time.Sleep(BlockRetryInterval)
				continue
			}

			// Parse out events
			err = l.getDepositEventsForBlock(latestBlock)
			if err != nil {
				l.log.Error("Failed to get events for block", "block", currentBlock, "err", err)
				retry--
				continue
			}

			retry = BlockRetryLimit
		}
	}
}

// getDepositEventsForBlock looks for the deposit event in the latest block
func (l *listener) getDepositEventsForBlock(latestBlock *big.Int) error {

	// resp, err := http.Get("https://api.shasta.trongrid.io/v1/contracts/TBzw9MxSMS7Eq7eUczhzjESGYsUR7UvQmA/events")
	// if err != nil {
	// 	return nil
	// }

	// if resp != nil {
	// 	body, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		return fmt.Errorf("parse contracts event error")
	// 	}
	// 	l.log.Debug("Contract Get Info", "Event", string(body))
	// }
	// l.log.Debug("Querying block for deposit events", "block", latestBlock)
	// query := buildQuery(l.cfg.bridgeContract, utils.MapTransferOut, latestBlock, latestBlock)

	// querying for logs
	// logs, err := l.conn.Client().FilterLogs(context.Background(), query)
	// if err != nil {
	// 	return fmt.Errorf("unable to Filter Logs: %w", err)
	// }

	// // read through the log events and handle their deposit event if handler is recognized
	// for _, log := range logs {
	// 	var m msg.Message
	// 	text := fmt.Sprintf("log out:%v", log)
	// 	l.log.Debug("Log for event ------>", "log", text)

	// 	err = l.router.Send(m)
	// 	if err != nil {
	// 		l.log.Error("subscription error: failed to route message", "err", err)
	// 	}
	// }

	return nil
}

// buildQuery constructs a query for the bridgeContract by hashing sig to get the event topic
func buildQuery(contract ethcommon.Address, sig EventSig, startBlock *big.Int, endBlock *big.Int) eth.FilterQuery {
	query := eth.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []ethcommon.Address{contract},
		// Topics: [][]ethcommon.Hash{
		// 	{sig.GetTopic()},
		// },
	}
	return query
}
