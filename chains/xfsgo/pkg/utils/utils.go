package utils

import (
	"fmt"
	"math/big"
	"strings"
)

const (
	decimalBase = 10
	hexBase     = 16
)

// valueToBig converts a string value to a *big.Int in the provided base
func valueToBig(value string, base int) (*big.Int, error) {
	val, ok := new(big.Int).SetString(value, base)
	if !ok {
		return nil, fmt.Errorf("unable to parse value")
	}

	return val, nil
}

// ParseUint256OrHex parses a string value as either a base 10 number
// or as a hex value
func ParseUint256OrHex(value *string) (*big.Int, error) {
	// Check if the value is valid
	if value == nil {
		return nil, fmt.Errorf("invalid value")
	}

	// Check if the value is hex
	if strings.HasPrefix(*value, "0x") {
		// Hex value, remove the prefix and parse it
		clipped := (*value)[2:]
		return valueToBig(clipped, hexBase)
	}

	// Decimal number, parse it
	return valueToBig(*value, decimalBase)
}
