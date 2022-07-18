package xfsgo

import (
	"fmt"
	"math/big"
	"strconv"

	"bridgeswap/controller/core"
)

var (
	HttpOpt       = "http"
	StartBlockOpt = "startblock"
)

type Config struct {
	name               string // Human-readable chain name
	id                 uint8  // ChainID
	endpoint           string // url for rpc endpoint
	http               bool   // Config for type of connection
	from               string // address of key to use
	keystorePath       string // Location of key files
	insecure           bool   // Indicated whether the test keyring should be used
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
		keystorePath:       chainCfg.KeystorePath,
		insecure:           chainCfg.Insecure,
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

func parseUseExtended(cfg *core.ChainConfig) bool {
	if b, ok := cfg.Opts["useExtendedCall"]; ok {
		res, err := strconv.ParseBool(b)
		if err != nil {
			panic(err)
		}
		return res
	}
	return false
}
