package tron

import "math/big"

type ContractData struct {
	Data []DataInfo `mapstructure:"data" json:"data" yaml:"data"`
}

type DataInfo struct {
	BlockNumer *big.Int          `mapstructure:"block_number" json:"block_number" yaml:"block_number"` // 块高度
	EventName  string            `mapstructure:"event_name" json:"event_name" yaml:"event_name"`       // 事件名称
	Result     map[string]string `mapstructure:"result" json:"result" yaml:"result"`                   // 事件结结果
}
