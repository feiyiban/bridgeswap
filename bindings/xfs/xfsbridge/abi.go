package xfsbridge

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"bridgeswap/sdk/xfs/common"
)

type Method struct {
	Name       string    `json:"name"`
	Argc       int       `json:"argc"`
	Args       Arguments `json:"args"`
	ReturnType string    `json:"returnType"`
}

type Events struct {
	Name string     `json:"name"`
	Argc int        `json:"argc"`
	Args ArgsEvents `json:"args"`
}

type ArgsEvent struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ArgsEvents []*ArgsEvent
type ABI struct {
	Events  map[string]*Events
	Methods map[string]*Method
}

func JSON(abi string) (ABI, error) {
	var jsonData ABI
	err := json.Unmarshal([]byte(abi), &jsonData)
	if err != nil {
		return ABI{}, nil
	}
	return jsonData, nil
}

func (abi ABI) PackArgs(name string, args ...interface{}) ([]byte, error) {
	method, exist := abi.Methods[name]
	if !exist {
		return nil, fmt.Errorf("method '%s' not found", name)
	}
	arguments, err := method.Args.Pack(args...)
	if err != nil {
		return nil, err
	}
	return arguments, nil
}

func (abi ABI) Create(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(CREATE, params...)
	if err != nil {
		return "", err
	}
	data := common.FromHex(BRIDGEBIN)
	dataMethod := common.FromHex(CREATE)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) TransferIn(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(TANSFERIN, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(BRIDGEBIN)
	dataMethod := common.FromHex(TANSFERIN)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) TransferOut(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(TANSFERINOUT, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(BRIDGEBIN)
	dataMethod := common.FromHex(TANSFERINOUT)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}
