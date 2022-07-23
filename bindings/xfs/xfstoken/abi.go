package xfstoken

import (
	"bridgeswap/sdk/xfsgo/common"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
)

type CTypeUint8 [1]byte
type CTypeBool [1]byte
type CTypeUint16 [2]byte
type CTypeUint32 [4]byte
type CTypeUint64 [8]byte
type CTypeUint256 [32]byte
type CTypeString []byte
type CTypeAddress [25]byte

var (
	CBoolTrue  = CTypeBool{1}
	CBoolFalse = CTypeBool{0}
)

func (t CTypeUint8) Uint8() uint8 {
	return t[0]
}

func (t CTypeUint8) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}
func (t *CTypeUint8) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	copy(t[:], bs)
	return
}

func (t CTypeUint16) Uint16() uint16 {
	return binary.LittleEndian.Uint16(t[:])
}

func (t CTypeUint16) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}
func (t *CTypeUint16) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	copy(t[:], bs)
	return
}
func (t CTypeUint32) Uint32() uint32 {
	return binary.LittleEndian.Uint32(t[:])
}
func (t CTypeUint32) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}
func (t *CTypeUint32) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	copy(t[:], bs)
	return
}
func (t CTypeUint64) Uint64() uint64 {
	return binary.LittleEndian.Uint64(t[:])
}
func (t CTypeUint64) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}
func (t *CTypeUint64) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	copy(t[:], bs)
	return
}
func (t CTypeUint256) BigInt() *big.Int {
	return new(big.Int).SetBytes(t[:])
}
func (t CTypeUint256) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}
func (t *CTypeUint256) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	copy(t[:], bs)
	return
}

func (t CTypeString) String() string {
	return string(t[:])
}

func (t CTypeString) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}
func (t *CTypeString) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	*t = make([]byte, len(bs))
	copy(*t, bs)
	return
}
func (t CTypeAddress) MarshalText() (d []byte, err error) {
	ds := hex.EncodeToString(t[:])
	d = make([]byte, len(ds))
	copy(d[:], ds[:])
	return
}

func (t *CTypeAddress) UnmarshalText(text []byte) (err error) {
	var bs []byte
	bs, err = hex.DecodeString(string(text))
	copy(t[:], bs)
	return
}
func (t CTypeAddress) Address() common.Address {
	return common.Bytes2Address(t[:])
}

func (t CTypeBool) Bool() bool {
	if t[0] == 1 {
		return true
	}
	return false
}

func NewUint8(n uint8) CTypeUint8 {
	return CTypeUint8{n}
}

func NewUint16(n uint16) (m CTypeUint16) {
	binary.LittleEndian.PutUint16(m[:], n)
	return
}

func NewUint32(n uint32) (m CTypeUint32) {
	binary.LittleEndian.PutUint32(m[:], n)
	return
}
func NewUint64(n uint64) (m CTypeUint64) {
	binary.LittleEndian.PutUint64(m[:], n)
	return
}

func NewUint256(n *big.Int) (m CTypeUint256) {
	bs := n.Bytes()
	r := len(m) - len(bs)
	copy(m[r:], bs)
	return
}

func NewAddress(n common.Address) (m CTypeAddress) {
	copy(m[:], n[:])
	return
}

type Method struct {
	Name       string    `json:"name"`
	Argc       int       `json:"argc"`
	Args       Arguments `json:"args"`
	ReturnType string    `json:"returnType"`
}

type ABI struct {
	EVENTS  map[string]*Method
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
	args, err := abi.PackArgs(CREATE, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(CREATE)
	data = append(data, dataMethod...)
	data = append(data, args...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) Name() (string, error) {
	args, err := abi.PackArgs(NAME)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(NAME)
	data = append(data, dataMethod...)
	data = append(data, args...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) Symbol() (string, error) {
	args, err := abi.PackArgs(SYMBOL)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(SYMBOL)
	data = append(data, dataMethod...)
	data = append(data, args...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) GetDecimals() (string, error) {
	args, err := abi.PackArgs(GETDECIMALS)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(GETDECIMALS)
	data = append(data, dataMethod...)
	data = append(data, args...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) GetTotalSupply() (string, error) {
	args, err := abi.PackArgs(GETTOTALSUPPLY)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(GETTOTALSUPPLY)
	data = append(data, dataMethod...)
	data = append(data, args...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) BalanceOf(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(BALANCEOF, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(BALANCEOF)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) Mint(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(MINT, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(MINT)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil

}

func (abi ABI) Burn(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(BURN, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(BURN)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) Approve(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(APPROVE, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(APPROVE)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) Allowance(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(ALLOWANCE, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(ALLOWANCE)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) TransferFrom(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(TRANSFERFROM, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(TRANSFERFROM)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}
