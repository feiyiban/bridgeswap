package xfstoken

import (
	"bridgeswap/sdk/xfsgo/common"
	"encoding/hex"
	"math/big"
	"testing"
)

type testToken struct {
	name        CTypeString
	symbol      CTypeString
	decimals    CTypeUint8
	totalSupply CTypeUint256
}

func TestDeployXFSToken(t *testing.T) {
	testACToken := testToken{
		name:        CTypeString("1"),
		symbol:      CTypeString("1"),
		decimals:    NewUint8(1),
		totalSupply: NewUint256(new(big.Int).SetInt64(1)),
	}
	// data := DeployXFSToken(testACToken.name, testACToken.symbol, testACToken.decimals, testACToken.totalSupply)
	abi, err := JSON(XFSTOKENABI)
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}
	packed, err := abi.PackArgs(CREATE, testACToken.name, testACToken.symbol, testACToken.decimals, testACToken.totalSupply)
	if err != nil {
		t.Errorf("PackArgs:%v\n", err.Error())
	}

	result := "0x" + hex.EncodeToString(packed)
	t.Logf("%v", result)
}

func TestMinerXFSToken(t *testing.T) {
	testACToken := testToken{
		name:        CTypeString("1"),
		symbol:      CTypeString("1"),
		decimals:    NewUint8(1),
		totalSupply: NewUint256(new(big.Int).SetInt64(1)),
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(CREATE)
	data = append(data, dataMethod...)
	// data := DeployXFSToken(testACToken.name, testACToken.symbol, testACToken.decimals, testACToken.totalSupply)
	abi, err := JSON(XFSTOKENABI)
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}
	packed, err := abi.PackArgs(CREATE, testACToken.name, testACToken.symbol, testACToken.decimals, testACToken.totalSupply)
	if err != nil {
		t.Errorf("PackArgs:%v\n", err.Error())
	}

	data = append(data, packed...)
	result := "0x" + hex.EncodeToString(data)
	t.Logf("%v", result)
}
