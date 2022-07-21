package events

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/common/ahash"
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

type EventABI struct {
	Events  map[string]*Events
	Methods map[string]*Method
}

func JSON(abi string) (EventABI, error) {
	var jsonData EventABI
	err := json.Unmarshal([]byte(abi), &jsonData)
	if err != nil {
		return EventABI{}, nil
	}
	return jsonData, nil
}

func (eventabi EventABI) packEventsName(hashevents string, events []*Event) ([]*Event, error) {
	for _, v := range eventabi.Events {
		eventHash := ahash.SHA256Array([]byte(v.Name))
		chash := common.Bytes2Hash(eventHash[:])
		if chash.Hex() == hashevents {
			return v.Args.Pack(events)
		}
	}
	return nil, fmt.Errorf("not events")
}

func (eventabi EventABI) Decode(hash, eventvalue string) (map[string]interface{}, error) {
	rawValue, err := hex.DecodeString(eventvalue)
	if err != nil {
		return nil, err
	}
	events, err := Str2Events(string(rawValue))
	if err != nil {
		return nil, err
	}
	eventsobj, err := eventabi.packEventsName(hash, events)
	if err != nil {
		return nil, err
	}
	return Events2Map(eventsobj), nil
}
