package ethereum

import (
	"crypto/ecdsa"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/ssh/terminal"

	"bridgeswap/bindings/eth/bridgev1"
	"bridgeswap/blockstore"
	"bridgeswap/chains/ethereum/connection"
	"bridgeswap/controller/core"
	"bridgeswap/logger"
	"bridgeswap/sdk/ethereum/store"

	"math/big"
)

type Connection interface {
	Connect() error
	GetPrivateKey() *ecdsa.PrivateKey
	Client() *ethclient.Client
	EnsureHasBytecode(addr common.Address) error
	LatestBlock() (*big.Int, error)
	Close()
}

type Chain struct {
	cfg      *core.ChainConfig // The config of the chain
	conn     Connection        // THe chains connection
	listener *listener         // The listener of this chain
	writer   *writer           // The writer of the chain
	stop     chan<- int
}

// getPassphrase fetches the correct passphrase depending on if a file is available to
// read from or if the user wants to enter in their own passphrase. Otherwise, just use
// the default passphrase. No confirmation of passphrase
func getPassphrase() (string, error) {
	fmt.Println("Enter password for tron key:")
	pass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	return string(pass), nil
}

func InitializeChain(chainCfg *core.ChainConfig, sysErr chan<- error, log logger.Logger) (*Chain, error) {
	cfg, err := parseChainConfig(chainCfg)
	if err != nil {
		return nil, err
	}

	passphrase, err := getPassphrase()
	if err != nil {
		return nil, err
	}

	ks, account, err := store.UnlockedKeystore(cfg.from, passphrase)
	if err != nil {
		return nil, err
	}

	_, key, err := ks.GetDecryptedKey(*account, passphrase)

	prikey := key.PrivateKey

	// cryptoKeyPair, err := keystore.KeypairFromAddress(cfg.from, keystore.EthChain, cfg.keystorePath)
	// if err != nil {
	// 	return nil, err
	// }
	// secp256k1KeyPair, ok := cryptoKeyPair.(*secp256k1.Keypair)
	// if !ok {
	// 	return nil, err
	// }

	blockStore, err := blockstore.NewBlockstore(cfg.blockstorePath, cfg.id, cfg.from)
	if err != nil {
		return nil, err
	}
	if !cfg.bFreshStart {
		latestBlock, err := blockStore.TryLoadLatestBlock()
		if err != nil {
			return nil, err
		}

		if latestBlock.Cmp(cfg.startBlock) == 1 {
			cfg.startBlock = latestBlock
		}
	}

	conn := connection.NewConnection(cfg.http, cfg.endpoint, prikey, cfg.gasLimit, cfg.maxGasPrice, cfg.minGasPrice, log)
	err = conn.Connect()
	if err != nil {
		return nil, err
	}

	err = conn.EnsureHasBytecode(common.HexToAddress(cfg.bridgeContract))
	if err != nil {
		return nil, err
	}

	if cfg.erc20Contract != "" {
		err = conn.EnsureHasBytecode(common.HexToAddress(cfg.erc20Contract))
		if err != nil {
			return nil, err
		}
	}

	bridge, err := bridgev1.NewBridgev1(common.HexToAddress(cfg.bridgeContract), conn.Client())
	if err != nil {
		log.Debug("error happened when call NewBridgev1", "error", err)
		return nil, err
	}

	if cfg.bLatestBlock {
		curr, err := conn.LatestBlock()
		if err != nil {
			return nil, err
		}
		cfg.startBlock = curr
	}

	stop := make(chan int)
	listener := NewListener(cfg, conn, blockStore, sysErr, stop, log)
	listener.setContracts(bridge)

	writer := NewWriter(cfg, conn, sysErr, stop, log)
	writer.setContract(bridge)

	return &Chain{
		cfg:      chainCfg,
		conn:     conn,
		writer:   writer,
		listener: listener,
		stop:     stop,
	}, nil

}

func (chain *Chain) SetRouter(router *core.Router) {
	chain.writer.log.Info("SetRouter")
	router.Listen(chain.cfg.ID, chain.writer)
	chain.listener.setRouter(router)
}

func (chain *Chain) Start() error {
	chain.writer.log.Debug("started chain...")
	err := chain.listener.start()
	if err != nil {
		return err
	}

	err = chain.writer.start()
	if err != nil {
		return err
	}

	chain.writer.log.Debug("Successfully started chain")

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
