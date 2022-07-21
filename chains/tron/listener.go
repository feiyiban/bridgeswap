package tron

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"net/url"
	"time"

	"bridgeswap/blockstore"
	"bridgeswap/controller/msg"

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
func NewListener(conn Connection, cfg *Config, log logger.Logger, bs blockstore.Blockstorer, stop <-chan int, sysErr chan<- error) *listener {
	return &listener{
		cfg:        *cfg,
		log:        log,
		conn:       conn,
		blockstore: bs,

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
func (listen *listener) setRouter(r chains.Router) {
	listen.router = r
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

			// if l.metrics != nil {
			// 	l.metrics.LatestKnownBlock.Set(float64(latestBlock.Int64()))
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
func (listen *listener) getDepositEventsForBlock(latestBlock *big.Int) error {
	listen.log.Debug("Querying block for deposit events", "block", latestBlock)

	params := url.Values{}
	parseURL, err := url.Parse(listen.cfg.erc20event)
	if err != nil {
		listen.log.Debug("Parse", listen.cfg.erc20event, err)
		return err
	}

	params.Set("block_number", latestBlock.String())
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	resp, err := http.Get(urlPathWithParams)
	defer resp.Body.Close()
	if err != nil {
		listen.log.Debug("Parse", listen.cfg.erc20event, err)
		return err
	}

	if resp == nil {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("parse contracts event error")
	}

	data := ContractData{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	for _, val := range data.Data {

		if val.EventName != "MAPTransferOut" {
			return nil
		}
		selfChainID := val.Result["fromChain"]
		destChainID := val.Result["toChain"]

		selfID, _ := big.NewInt(0).SetString(selfChainID, 10)

		destID, _ := big.NewInt(0).SetString(destChainID, 10)

		listen.log.Debug("Log for event ------>", "selfChainID", selfChainID, "destChainID", destChainID)

		byteTo := []byte(val.Result["to"])
		byteValue := []byte(val.Result["amount"])

		m := msg.Message{
			Source:      uint8(selfID.Uint64()),
			Destination: uint8(destID.Uint64()),
			Type:        msg.TokenTransfer,
			Payload: []interface{}{
				string(byteTo),
				string(byteValue),
			},
		}

		listen.log.Debug("Log for event ------>", "Payload", m.Payload)

		if err != nil {
			return err
		}

		err = listen.router.Send(m)
		if err != nil {
			listen.log.Error("subscription error: failed to route message", "err", err)
		}
	}

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
