package types

import (
	"bridgeswap/sdk/xfsgo/common"
	"bridgeswap/sdk/xfsgo/common/ahash"
	"bridgeswap/sdk/xfsgo/crypto"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
)

type BlockHeader struct {
	Height        uint64         `json:"height"`
	Version       uint32         `json:"version"`
	HashPrevBlock common.Hash    `json:"hash_prev_block"`
	Timestamp     uint64         `json:"timestamp"`
	Coinbase      common.Address `json:"coinbase"`
	// merkle tree root hash
	StateRoot        common.Hash `json:"state_root"`
	TransactionsRoot common.Hash `json:"transactions_root"`
	ReceiptsRoot     common.Hash `json:"receipts_root"`
	GasLimit         *big.Int    `json:"gas_limit"`
	GasUsed          *big.Int    `json:"gas_used"`
	// pow consensus.
	Bits       uint32 `json:"bits"`
	Nonce      uint32 `json:"nonce"`
	ExtraNonce uint64 `json:"extranonce"`
}

type StringRawTransaction struct {
	Version   string `json:"version"`
	From      string `json:"from"`
	To        string `json:"to"`
	Value     string `json:"value"`
	Data      string `json:"data"`
	GasLimit  string `json:"gas_limit"`
	GasPrice  string `json:"gas_price"`
	Signature string `json:"signature"`
	Nonce     string `json:"nonce"`
}

func (tx *StringRawTransaction) SignWithPrivateKey(fromprikey string) error {

	keyEnc := fromprikey
	if keyEnc[0] == '0' && keyEnc[1] == 'x' {
		keyEnc = keyEnc[2:]
	} else {
		return errors.New("binary forward backward error")
	}

	keyDer, err := hex.DecodeString(keyEnc)
	if err != nil {
		return err
	}

	_, key, err := crypto.DecodePrivateKey(keyDer)
	if err != nil {
		return err
	}

	hash := tx.SignHash()
	sig, err := crypto.ECDSASign(hash.Bytes(), key)
	if err != nil {
		return err
	}
	tx.Signature = hex.EncodeToString(sig)
	return nil
}

// Transfer2Raw trading partner code Base64 format
func (tx *StringRawTransaction) Transfer2Raw() (string, error) {
	bs, err := json.Marshal(tx)
	if err != nil {
		return "", err
	}
	result := base64.StdEncoding.EncodeToString(bs)
	return result, nil
}

// signHash generate transaction hash
func (tx *StringRawTransaction) SignHash() common.Hash {
	tmp := map[string]string{
		"version":   tx.Version,
		"to":        tx.To,
		"gas_price": tx.GasPrice,
		"gas_limit": tx.GasLimit,
		"data":      tx.Data,
		"nonce":     tx.Nonce,
		"value":     tx.Value,
	}
	enc := common.SortAndEncodeMap(tmp)
	if enc == "" {
		return common.Hash{}
	}
	return common.Bytes2Hash(ahash.SHA256([]byte(enc)))
}

type Receipt struct {
	// Consensus fields: These fields are defined by the Yellow Paper
	Type              uint8  `json:"type,omitempty"`
	Status            uint64 `json:"status"`
	CumulativeGasUsed uint64 `json:"cumulativeGasUsed" gencodec:"required"`

	// Implementation fields: These fields are added by geth when processing a transaction.
	// They are stored in the chain database.
	TxHash          string `json:"transactionHash" gencodec:"required"`
	ContractAddress string `json:"contractAddress"`
	GasUsed         uint64 `json:"gasUsed" gencodec:"required"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash        string   `json:"blockHash,omitempty"`
	BlockNumber      *big.Int `json:"blockNumber,omitempty"`
	TransactionIndex uint     `json:"transactionIndex"`
}
