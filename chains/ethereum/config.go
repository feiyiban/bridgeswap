package ethereum

import (
	"errors"
	"fmt"
	"math/big"

	"bridgeswap/controller/core"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const DefaultKeystorePath = "./keys"

const (
	DefaultGasLimit           = 6721975
	DefaultGasPrice           = 20000000000
	DefaultMinGasPrice        = 0
	DefaultBlockConfirmations = 3
	DefaultGasMultiplier      = 1
)

var (
	BridgeOpt             = "bridge"
	Erc20HandlerOpt       = "erc20Token"
	MaxGasPriceOpt        = "maxGasPrice"
	MinGasPriceOpt        = "minGasPrice"
	GasLimitOpt           = "gasLimit"
	GasMultiplier         = "gasMultiplier"
	HttpOpt               = "http"
	StartBlockOpt         = "startblock"
	BlockConfirmationsOpt = "blockConfirmations"
)

// Config encapsulates all necessary parameters in ethereum compatible forms
type Config struct {
	name               string // Human-readable chain name
	id                 uint8  // ChainID
	endpoint           string // url for rpc endpoint
	from               string // address of key to use
	keystorePath       string // Location of keyfiles
	blockstorePath     string
	freshStart         bool // Disables loading from blockstore at start
	bridgeContract     common.Address
	erc20Contract      common.Address
	gasLimit           *big.Int
	maxGasPrice        *big.Int
	minGasPrice        *big.Int
	gasMultiplier      *big.Float
	http               bool // Config for type of connection
	startBlock         *big.Int
	blockConfirmations *big.Int
	egsApiKey          string // API key for ethgasstation to query gas prices
	egsSpeed           string // The speed which a transaction should be processed: average, fast, fastest. Default: fast
}

// parseChainConfig uses a core.ChainConfig to construct a corresponding Config
func parseChainConfig(chainCfg *core.ChainConfig) (*Config, error) {

	config := &Config{
		name:               chainCfg.Name,
		id:                 chainCfg.ID,
		endpoint:           chainCfg.Endpoint,
		from:               chainCfg.From,
		keystorePath:       chainCfg.KeystorePath,
		blockstorePath:     chainCfg.BlockstorePath,
		freshStart:         chainCfg.FreshStart,
		bridgeContract:     common.Address{},
		erc20Contract:      common.Address{},
		gasLimit:           big.NewInt(DefaultGasLimit),
		maxGasPrice:        big.NewInt(DefaultGasPrice),
		minGasPrice:        big.NewInt(DefaultMinGasPrice),
		gasMultiplier:      big.NewFloat(DefaultGasMultiplier),
		http:               false,
		startBlock:         big.NewInt(0),
		blockConfirmations: big.NewInt(0),
		egsApiKey:          "",
		egsSpeed:           "",
	}

	if contract, ok := chainCfg.Opts[BridgeOpt]; ok && contract != "" {
		config.bridgeContract = common.HexToAddress(contract)
		delete(chainCfg.Opts, BridgeOpt)
	} else {
		return nil, fmt.Errorf("must provide opts.bridge field for ethereum config")
	}

	if contract, ok := chainCfg.Opts[Erc20HandlerOpt]; ok {
		config.erc20Contract = common.HexToAddress(contract)
		delete(chainCfg.Opts, Erc20HandlerOpt)
	}

	if gasPrice, ok := chainCfg.Opts[MaxGasPriceOpt]; ok {
		price, parseErr := hexutil.DecodeBig(gasPrice)
		if parseErr != nil {
			return nil, fmt.Errorf("unable to parse max gas price, %w", parseErr)
		}

		config.maxGasPrice = price
		delete(chainCfg.Opts, MaxGasPriceOpt)
	}

	if minGasPrice, ok := chainCfg.Opts[MinGasPriceOpt]; ok {

		price, parseErr := hexutil.DecodeBig(minGasPrice)
		if parseErr != nil {
			return nil, fmt.Errorf("unable to parse min gas price, %w", parseErr)
		}

		config.minGasPrice = price
		delete(chainCfg.Opts, MinGasPriceOpt)
	}

	if gasLimit, ok := chainCfg.Opts[GasLimitOpt]; ok {
		limit, parseErr := hexutil.DecodeBig(gasLimit)
		if parseErr != nil {
			return nil, fmt.Errorf("unable to parse gas limit, %w", parseErr)
		}

		config.gasLimit = limit
		delete(chainCfg.Opts, GasLimitOpt)
	}

	if gasMultiplier, ok := chainCfg.Opts[GasMultiplier]; ok {
		multilier := big.NewFloat(1)
		_, pass := multilier.SetString(gasMultiplier)
		if pass {
			config.gasMultiplier = multilier
			delete(chainCfg.Opts, GasMultiplier)
		} else {
			return nil, errors.New("unable to parse gasMultiplier to float")
		}
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
