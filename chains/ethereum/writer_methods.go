package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (write *writer) BridgeTransferIn(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int, _fromChain *big.Int, _toChain *big.Int) error {
	tx, err := write.bridge.TransferIn(opts, _token, _to, _amount, _fromChain, _toChain)
	if err != nil {
		write.log.Error("Failed to TransferIn by Bridge", "err", err)
		return err
	}

	write.log.Debug("Bridge", "TransferIn", tx)
	return nil
}
