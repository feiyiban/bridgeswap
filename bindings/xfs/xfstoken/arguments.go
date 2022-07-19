package xfstoken

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

func writeStringParams(w Buffer, s CTypeString) {
	slen := len(s)
	var slenbuf [8]byte
	binary.LittleEndian.PutUint64(slenbuf[:], uint64(slen))
	_, _ = w.Write(slenbuf[:])
	_, _ = w.Write(s)
}

type ArgObj struct {
	Type string `json:"type"`
}

type Arguments []*ArgObj

func (arguments Arguments) Pack(args ...interface{}) ([]byte, error) {
	// Make sure arguments match up and pack them
	abiArgs := arguments
	if len(args) != len(abiArgs) {
		return nil, fmt.Errorf("argument count mismatch: got %d for %d", len(args), len(abiArgs))
	}

	buf := NewBuffer(nil)
	for offset := int(0); offset < len(arguments); offset++ {
		argObj := arguments[offset]
		argValue := args[offset]

		ObjType := argObj.Type

		if ObjType == "CTypeString" {
			typeofValue := reflect.TypeOf(argValue)
			if typeofValue.Name() != "CTypeString" {
				return nil, fmt.Errorf("param type check err")
			}
			para := argValue.(CTypeString)
			writeStringParams(buf, para)
		} else if ObjType == "CTypeUint8" {
			typeofValue := reflect.TypeOf(argValue)
			if typeofValue.Name() != "CTypeUint8" {
				return nil, fmt.Errorf("param type check err")
			}
			para := argValue.(CTypeUint8)
			buf.Write(para[:])
		} else if ObjType == "CTypeUint256" {
			typeofValue := reflect.TypeOf(argValue)
			if typeofValue.Name() != "CTypeUint256" {
				return nil, fmt.Errorf("param type check err")
			}
			para := argValue.(CTypeUint256)
			buf.Write(para[:])
		} else if ObjType == "CTypeAddress" {
			typeofValue := reflect.TypeOf(argValue)
			if typeofValue.Name() != "CTypeAddress" {
				return nil, fmt.Errorf("param type check err")
			}
			para := argValue.(CTypeAddress)
			buf.Write(para[:])
		} else {
			return nil, fmt.Errorf("type check err")
		}
	}

	return buf.Bytes(), nil
}
