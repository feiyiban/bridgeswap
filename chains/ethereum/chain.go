package ethereum

import (
	"bridgeswap/controller/core"
	"bridgeswap/logger"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Connection interface {
	Connect() error
	Opts() *bind.TransactOpts
	CallOpts() *bind.CallOpts
	LockAndUpdateOpts() error
	UnlockOpts()
	Client() *ethclient.Client
	EnsureHasBytecode(address common.Address) error
	LatestBlock() (*big.Int, error)
	WaitForBlock(block *big.Int, delay *big.Int) error
	Close()
}

type Chain struct {
	cfg  *core.ChainConfig // The config of the chain
	conn Connection        // THe chains connection
	stop chan<- int
}

func InitializeChain(chainCfg *core.ChainConfig, log logger.Logger, sysErr chan<- error) (*Chain, error) {
	return nil, nil
}

func (chain *Chain) SetRouter(router *core.Router) {

}

func (chain *Chain) Start() error {
	return nil
}

func (chain *Chain) ID() uint8 {
	return chain.cfg.ID
}

func (chain *Chain) Name() string {
	return chain.cfg.Name
}

func (chain *Chain) Stop() {
	close(chain.stop)
	if chain.conn != nil {
		chain.conn.Close()
	}
}
