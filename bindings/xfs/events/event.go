package events

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/core/abi"
)

type Event struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

func Hex(events []*Event) (string, error) {
	bs, err := json.Marshal(events)
	return string(bs), err
}

func Map2Hex(events []*Event) map[string]interface{} {
	res := make(map[string]interface{})
	for _, event := range events {
		res[event.Name] = event.Value
	}
	return res
}

func Str2Events(jsonEvents string) ([]*Event, error) {
	events := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonEvents), &events)
	if err != nil {
		return nil, err
	}
	es := Map2Events(events)
	return es, nil
}

func Events2Map(events []*Event) map[string]interface{} {
	result := make(map[string]interface{})
	for _, v := range events {
		result[v.Name] = v.Value
	}
	return result
}

func Map2Events(mapEvents map[string]interface{}) []*Event {
	events := make([]*Event, 0)
	for name, value := range mapEvents {
		event := new(Event)
		event.Name = name
		event.Value = value
		events = append(events, event)
	}
	return events
}

type ArgsEvent struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ArgsEvents []*ArgsEvent

func (argevents ArgsEvents) Pack(args []*Event) ([]*Event, error) {
	// Make sure argevents match up and pack them
	abiArgs := argevents
	if len(args) != len(abiArgs) {
		return nil, fmt.Errorf("argument count mismatch: got %d for %d", len(args), len(abiArgs))
	}

	events := make([]*Event, 0)
	for _, v := range args {
		event := new(Event)
		event.Name = v.Name
		vas := v.Value.(string)
		for _, obj := range argevents {

			if strings.EqualFold(strings.ToLower(obj.Name), strings.ToLower(v.Name)) {
				if obj.Type == "CTypeString" {
					para := abi.CTypeString(vas)
					pars, err := hex.DecodeString(para.String())
					if err != nil {
						return nil, fmt.Errorf("ctypestring to string:%v", v.Value)
					}
					event.Value = string(pars)
				} else if obj.Type == "CTypeUint8" {
					big8, ok := big.NewInt(0).SetString(vas, 16)
					if !ok {
						return nil, fmt.Errorf("ctypeuint256 to uint256:%v", v.Value)
					}
					para := abi.NewUint8(uint8(big8.Uint64()))
					event.Value = para.Uint8()
				} else if obj.Type == "CTypeUint256" {
					big256, ok := big.NewInt(0).SetString(vas, 16)
					if !ok {
						return nil, fmt.Errorf("ctypeuint256 to uint256:%v", v.Value)
					}
					para := abi.NewUint256(big256)
					event.Value = para.BigInt().String()
				} else if obj.Type == "CTypeAddress" {
					addrBytes, err := hex.DecodeString(vas)
					if err != nil {
						return nil, fmt.Errorf("ctypeaddress to address:%v", err)
					}
					para := abi.NewAddress(common.Bytes2Address(addrBytes))
					address := para.Address()
					event.Value = address.B58String()
				} else if obj.Type == "CTypeBool" {
					para := v.Value.(abi.CTypeBool)
					event.Value = para.Bool()
				} else {
					return nil, fmt.Errorf("type check err")
				}
			}
		}
		events = append(events, event)
	}
	return events, nil
}
