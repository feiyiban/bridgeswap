package xfsgo

import (
	"fmt"
	"math/big"

	"bridgeswap/controller/core"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	DefaultGasLimit           = 6721975
	DefaultGasPrice           = 20000000000
	DefaultMinGasPrice        = 0
	DefaultBlockConfirmations = 3
	DefaultGasMultiplier      = 1
)

var (
	BridgeOpt             = "bridge"
	ERC20HandlerOpt       = "erc20Token"
	StartBlockOpt         = "startBlock"
	BlockConfirmationsOpt = "blockConfirmations"
	MaxGasPriceOpt        = "maxGasPrice"
	MinGasPriceOpt        = "minGasPrice"
	GasLimitOpt           = "gasLimit"
	HttpOpt               = "http"
)

type Config struct {
	name               string   // Human-readable chain name
	id                 uint8    // Human-readable chain id
	endpoint           string   // url for rpc endpoint
	from               string   // address of key to use
	keystorePath       string   // Location of keyfiles
	blockstorePath     string   // Synchronous block record path
	bFreshStart        bool     // Disables loading from blockstore at start
	bLatestBlock       bool     // If true, overrides blockstore or latest block in config and starts from current block
	http               bool     // Config for type of connection
	bridgeContract     string   // bridge contract address
	erc20Contract      string   // erc20 token contract address
	gasLimit           *big.Int // gas limit
	maxGasPrice        *big.Int // max gas price
	minGasPrice        *big.Int // min gas price
	startBlock         *big.Int // Start synchronizing block height
	blockConfirmations *big.Int // Synchronize the block to the specified offset post-processing block
}

// parseChainConfig uses a contraoller.core.ChainConfig to construct a corresponding Config
// The generated configuration file is also used for listener and writer
func parseChainConfig(chainCfg *core.ChainConfig) (*Config, error) {
	config := &Config{
		name:               chainCfg.Name,
		id:                 chainCfg.ID,
		endpoint:           chainCfg.Endpoint,
		from:               chainCfg.From,
		keystorePath:       chainCfg.KeystorePath,
		blockstorePath:     chainCfg.BlockstorePath,
		bFreshStart:        chainCfg.BFreshStart,
		bLatestBlock:       chainCfg.BLatestBlock,
		http:               false,
		bridgeContract:     "",
		erc20Contract:      "",
		gasLimit:           big.NewInt(DefaultGasLimit),
		maxGasPrice:        big.NewInt(DefaultGasPrice),
		minGasPrice:        big.NewInt(DefaultMinGasPrice),
		startBlock:         big.NewInt(0),
		blockConfirmations: big.NewInt(0),
	}

	if bridgeContract, ok := chainCfg.Opts[BridgeOpt]; ok && bridgeContract != "" {
		config.bridgeContract = bridgeContract
		delete(chainCfg.Opts, BridgeOpt)
	} else {
		return nil, fmt.Errorf("must provide opts.bridge field for ethereum config")
	}

	if erc20TokenContract, ok := chainCfg.Opts[ERC20HandlerOpt]; ok {
		config.erc20Contract = erc20TokenContract
		delete(chainCfg.Opts, ERC20HandlerOpt)
	}

	if maxGasPrice, ok := chainCfg.Opts[MaxGasPriceOpt]; ok {
		bigMaxGasPrice, parseErr := hexutil.DecodeBig(maxGasPrice)
		if parseErr != nil {
			return nil, fmt.Errorf("unable to parse max gas price, %v", parseErr)
		}
		config.maxGasPrice = bigMaxGasPrice
		delete(chainCfg.Opts, MaxGasPriceOpt)
	}

	if minGasPrice, ok := chainCfg.Opts[MinGasPriceOpt]; ok {
		bigMinGasPrice, parseErr := hexutil.DecodeBig(minGasPrice)
		if parseErr != nil {
			return nil, fmt.Errorf("unable to parse min gas price, %v", parseErr)
		}
		config.minGasPrice = bigMinGasPrice
		delete(chainCfg.Opts, MinGasPriceOpt)
	}

	if gasLimit, ok := chainCfg.Opts[GasLimitOpt]; ok {
		bigGaslimit, parseErr := hexutil.DecodeBig(gasLimit)
		if parseErr != nil {
			return nil, fmt.Errorf("unable to parse gas limit, %v", parseErr)
		}
		config.gasLimit = bigGaslimit
		delete(chainCfg.Opts, GasLimitOpt)
	}

	if bHttp, ok := chainCfg.Opts[HttpOpt]; ok && bHttp == "true" {
		config.http = true
		delete(chainCfg.Opts, HttpOpt)
	} else if HTTP, ok := chainCfg.Opts[HttpOpt]; ok && HTTP == "false" {
		config.http = false
		delete(chainCfg.Opts, HttpOpt)
	}

	if startBlock, ok := chainCfg.Opts[StartBlockOpt]; ok && startBlock != "" {
		bigStartBlock := big.NewInt(int64(0))
		startBlock, bPass := bigStartBlock.SetString(startBlock, 10)
		if bPass {
			config.startBlock = startBlock
			delete(chainCfg.Opts, StartBlockOpt)
		} else {
			return nil, fmt.Errorf("unable to parse %s", StartBlockOpt)
		}
	}

	if blockConfirmations, ok := chainCfg.Opts[BlockConfirmationsOpt]; ok && blockConfirmations != "" {
		bigBlockConfirm := big.NewInt(int64(DefaultBlockConfirmations))
		_, bPass := bigBlockConfirm.SetString(blockConfirmations, 10)
		if bPass {
			config.blockConfirmations = bigBlockConfirm
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
