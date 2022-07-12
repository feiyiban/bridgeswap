package ethereum

import (
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/crypto"
)

type EventSig string

func (es EventSig) GetTopic() common.Hash {
	return crypto.Keccak256Hash([]byte(es))
}

const (
	MapTransferOut EventSig = "MAPTransferOut(uint,uint,string,uint)"
)
