package types

import (
	"bridgeswap/chains/xfsgo/pkg/common"
	"math/big"
)

type GetLogsRequest struct {
	FromBlock string `json:"from_block"`
	ToBlock   string `json:"to_block"`
	Address   string `json:"address"`
	TxHash    string `json:"tx_hash"`
	EventHash string `json:"event_hash"`
}

type EventLogResp struct {
	BlockHeight     uint64         `json:"block_number"`
	BlockHash       common.Hash    `json:"block_hash"`
	TransactionHash common.Hash    `json:"transaction_hash"`
	EventHash       common.Hash    `json:"event_hash"`
	EventValue      string         `json:"event_value"`
	Address         common.Address `json:"address"`
}

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
