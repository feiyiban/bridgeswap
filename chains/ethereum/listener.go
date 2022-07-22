package ethereum

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"

	"bridgeswap/bindings/eth/bridgev1"
	"bridgeswap/blockstore"
	"bridgeswap/chains"
	"bridgeswap/controller/msg"
	"bridgeswap/logger"
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

			// Goto next block and reset retry counter
			currentBlock.Add(currentBlock, big.NewInt(1))
			retry = BlockRetryLimit
		}
	}
}

// buildQuery constructs a query for the bridgeContract by hashing sig to get the event topic
func buildQuery(contract common.Address, sig EventSig, startBlock *big.Int, endBlock *big.Int) ethereum.FilterQuery {
	query := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: []common.Address{contract},
		// Topics: [][]common.Hash{
		// 	{sig.GetTopic()},
		// },
	}
	return query
}

// getDepositEventsForBlock looks for the deposit event in the latest block
func (listen *listener) getDepositEventsForBlock(latestBlock *big.Int) error {
	listen.log.Debug("Querying block for deposit events", "block", latestBlock)

	query := buildQuery(listen.cfg.bridgeContract, MapTransferOut, latestBlock, latestBlock)

	// querying for logs
	logs, err := listen.conn.Client().FilterLogs(context.Background(), query)
	if err != nil {
		return fmt.Errorf("unable to Filter Logs: %w", err)
	}

	// read through the log events and handle their deposit event if handler is recognized
	for _, log := range logs {
		text := fmt.Sprintf("log out:%v", log)
		listen.log.Debug("Log for event ------>", "log", text)
		// l.log.Debug("Log for event ------>", "data", log.Data)

		if len(log.Topics) != 3 {
			return nil
		}
		selfChainID := uint8(log.Topics[1].Big().Uint64())
		destChainID := uint8(log.Topics[2].Big().Uint64())

		listen.log.Debug("Log for event ------>", "selfChainID", selfChainID, "destChainID", destChainID)

		if len(log.Data) <= 64 {
			listen.log.Debug("getDepositEventsForBlock ", "tranferIn", "not deal")
			return nil
		}

		bufferValue := new(bytes.Buffer)
		bufferValue.Write(log.Data[32:64])
		byteValue := bufferValue.Bytes()
		listen.log.Info("TransferOut--------->", "byteValue", byteValue)

		bufferAddrLenght := new(bytes.Buffer)
		bufferAddrLenght.Write(log.Data[64:96])
		byteAddrLength := bufferAddrLenght.Bytes()
		listen.log.Info("TransferOut--------->", "byteAddrLength", byteAddrLength)

		intLenght := big.NewInt(0).SetBytes(byteAddrLength)
		if err != nil {
			listen.log.Debug("getDepositEventsForBlock ", "strconv.Atoi", err)
			return nil
		}
		bufferAddr := new(bytes.Buffer)
		bufferAddr.Write(log.Data[96 : 96+intLenght.Int64()])

		byteAddr := bufferAddr.Bytes()
		listen.log.Info("TransferOut--------->", "toaddr", byteAddr)
		m := msg.Message{
			Source:      selfChainID,
			Destination: destChainID,
			Type:        msg.TokenTransfer,
			Payload: []interface{}{
				string(byteAddr),
				big.NewInt(0).SetBytes(byteValue).String(),
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
