package common

import (
	"errors"
	"os"
)

const (
	// DefaultConfigAccountAliasesDirName fro accounts
	DefaultConfigAccountAliasesDirName = "account-keys"
	// DefaultPassphrase for accounts
	DefaultPassphrase = ""
	// Secp256k1PrivateKeyBytesLength privete key
	Secp256k1PrivateKeyBytesLength = 34
	// AmountDecimalPoint TRX decimal point
	AmountDecimalPoint = 6
)

var (
	// DefaultConfigDirName for wallets
	DefaultConfigDirName      = ".xfsgobridge"
	DebugGRPC                 = false
	DebugTransaction          = false
	ErrNotAbsPath             = errors.New("keypath is not absolute path")
	ErrBadKeyLength           = errors.New("Invalid private key (wrong length)")
	ErrParseBytesOfPrivateKey = errors.New("can not parse the byte[] of private key")
	ErrFoundNoPass            = errors.New("found no passphrase file")
)

func init() {
	if _, enabled := os.LookupEnv("XFSGOCTL_GRPC_DEBUG"); enabled != false {
		DebugGRPC = true
	}
	if _, enabled := os.LookupEnv("XFSGOCTL_TX_DEBUG"); enabled != false {
		DebugTransaction = true
	}
	if _, enabled := os.LookupEnv("XFSGOCTL_ALL_DEBUG"); enabled != false {
		EnableAllVerbose()
	}
}

// EnableAllVerbose sets debug vars to true
func EnableAllVerbose() {
	DebugGRPC = true
	DebugTransaction = true
}
