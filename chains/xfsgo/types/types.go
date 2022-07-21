package types

import "bridgeswap/sdk/xfs/common"

type RawTransactionArgs struct {
	Data string `json:"data"`
}

type GetTranByHashArgs struct {
	Hash string `json:"hash"`
}

type GetAddrNonceByHashArgs struct {
	Address string `json:"address"`
}

type RemoveTxHashArgs struct {
	Hash string `json:"hash"`
}

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
