package xfsgo

import (
	"fmt"
	"math/big"
	"strconv"

	"bridgeswap/controller/core"
)

const (
	DefaultGasLimit           = 6721975
	DefaultGasPrice           = 20000000000
	DefaultMinGasPrice        = 0
	DefaultBlockConfirmations = 2
	DefaultGasMultiplier      = 1
)

var (
	HttpOpt               = "http"
	StartBlockOpt         = "startblock"
	BridgeOpt             = "bridge"
	Erc20HandlerOpt       = "erc20"
	BlockConfirmationsOpt = "blockConfirmations"
)

type Config struct {
	name               string // Human-readable chain name
	id                 uint8  // ChainID
	endpoint           string // url for rpc endpoint
	http               bool   // Config for type of connection
	from               string // address of key to use
	bridgeContract     string
	erc20Contract      string
	keystorePath       string // Location of key files
	blockstorePath     string // Location of blockstore
	freshStart         bool   // If true, blockstore is ignored at start.
	startBlock         *big.Int
	blockConfirmations *big.Int
	latestBlock        bool              // If true, overrides blockstore or latest block in config and starts from current block
	opts               map[string]string // Per chain options
}

func parseChainConfig(chainCfg *core.ChainConfig) (*Config, error) {

	config := &Config{
		name:               chainCfg.Name,
		id:                 chainCfg.ID,
		endpoint:           chainCfg.Endpoint,
		http:               false,
		startBlock:         big.NewInt(0),
		blockConfirmations: big.NewInt(0),
		from:               chainCfg.From,
		bridgeContract:     "",
		erc20Contract:      "",
		keystorePath:       chainCfg.KeystorePath,
		blockstorePath:     chainCfg.BlockstorePath,
		freshStart:         chainCfg.FreshStart,
		latestBlock:        chainCfg.LatestBlock,
		opts:               chainCfg.Opts,
	}

	if HTTP, ok := chainCfg.Opts[HttpOpt]; ok && HTTP == "true" {
		config.http = true
		delete(chainCfg.Opts, HttpOpt)
	} else if HTTP, ok := chainCfg.Opts[HttpOpt]; ok && HTTP == "false" {
		config.http = false
		delete(chainCfg.Opts, HttpOpt)
	}

	if startBlock, ok := chainCfg.Opts[StartBlockOpt]; ok && startBlock != "" {
		block := big.NewInt(0)
		startBlock, pass := block.SetString(startBlock, 10)
		if pass {
			config.startBlock = startBlock
			delete(chainCfg.Opts, StartBlockOpt)
		} else {
			return nil, fmt.Errorf("unable to parse %s", StartBlockOpt)
		}
	}

	if contract, ok := chainCfg.Opts[BridgeOpt]; ok && contract != "" {
		config.bridgeContract = contract
		delete(chainCfg.Opts, BridgeOpt)
	} else {
		return nil, fmt.Errorf("must provide opts.bridge field for ethereum config")
	}

	if contract, ok := chainCfg.Opts[Erc20HandlerOpt]; ok {
		config.erc20Contract = contract
		delete(chainCfg.Opts, Erc20HandlerOpt)
	}

	if blockConfirmations, ok := chainCfg.Opts[BlockConfirmationsOpt]; ok && blockConfirmations != "" {
		val := big.NewInt(DefaultBlockConfirmations)
		_, pass := val.SetString(blockConfirmations, 10)
		if pass {
			config.blockConfirmations = val
			delete(chainCfg.Opts, BlockConfirmationsOpt)
		} else {
			return nil, fmt.Errorf("unable to parse %s", BlockConfirmationsOpt)
		}
	} else {
		config.blockConfirmations = big.NewInt(DefaultBlockConfirmations)
		delete(chainCfg.Opts, BlockConfirmationsOpt)
	}

	return config, nil
}

func parseStartBlock(cfg *core.ChainConfig) uint64 {
	if blk, ok := cfg.Opts["startBlock"]; ok {
		res, err := strconv.ParseUint(blk, 10, 32)
		if err != nil {
			panic(err)
		}
		return res
	}
	return 0
}
