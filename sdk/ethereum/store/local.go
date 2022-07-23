package store

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"bridgeswap/sdk/ethereum/keystore"
	"bridgeswap/sdk/ethereum/sdkcommon"

	"github.com/ethereum/go-ethereum/common"

	"github.com/pkg/errors"

	homedir "github.com/mitchellh/go-homedir"
)

func init() {
	uDir, _ := homedir.Dir()
	tronCTLDir := path.Join(uDir, sdkcommon.DefaultConfigDirName, sdkcommon.DefaultConfigAccountAliasesDirName)
	if _, err := os.Stat(tronCTLDir); os.IsNotExist(err) {
		os.MkdirAll(tronCTLDir, 0700)
	}
}

// LocalAccounts returns a slice of local account alias names
func LocalAccounts() []string {
	uDir, _ := homedir.Dir()
	files, _ := ioutil.ReadDir(path.Join(
		uDir,
		sdkcommon.DefaultConfigDirName,
		sdkcommon.DefaultConfigAccountAliasesDirName,
	))
	accounts := []string{}

	for _, node := range files {
		if node.IsDir() {
			accounts = append(accounts, path.Base(node.Name()))
		}
	}
	return accounts
}

var (
	describe = fmt.Sprintf("%-24s\t\t%23s\n", "NAME", "ADDRESS")
	// ErrNoUnlockBadPassphrase for bad password
	ErrNoUnlockBadPassphrase = fmt.Errorf("could not unlock account with passphrase, perhaps need different phrase")
)

// DescribeLocalAccounts will display all the account alias name and their corresponding one address
func DescribeLocalAccounts() {
	fmt.Println(describe)
	for _, name := range LocalAccounts() {
		ks := FromAccountName(name)
		allAccounts := ks.Accounts()
		for _, account := range allAccounts {
			fmt.Printf("%-48s\t%s\n", name, account.Address)
		}
	}
}

// DoesNamedAccountExist return true if the given string name is an alias account already define,
// and return false otherwise
func DoesNamedAccountExist(name string) bool {
	for _, account := range LocalAccounts() {
		if account == name {
			return true
		}
	}
	return false
}

// AddressFromAccountName Returns address for account name if exists
func AddressFromAccountName(name string) (string, error) {
	ks := FromAccountName(name)
	// FIXME: Assume 1 account per keystore for now
	for _, account := range ks.Accounts() {
		return account.Address.Hex(), nil
	}
	return "", fmt.Errorf("keystore not found")
}

// FromAddress will return nil if the Base58 string is not found in the imported accounts
func FromAddress(addr string) *keystore.KeyStore {
	for _, name := range LocalAccounts() {
		ks := FromAccountName(name)
		allAccounts := ks.Accounts()
		for _, account := range allAccounts {
			if addr == account.Address.String() {
				return ks
			}
		}
	}
	return nil
}

// FromAccountName get account from name
func FromAccountName(name string) *keystore.KeyStore {
	uDir, _ := homedir.Dir()
	p := path.Join(uDir, sdkcommon.DefaultConfigDirName, sdkcommon.DefaultConfigAccountAliasesDirName, name)
	return keystore.ForPath(p)
}

// DefaultLocation get deafault location
func DefaultLocation() string {
	uDir, _ := homedir.Dir()
	return path.Join(uDir, sdkcommon.DefaultConfigDirName, sdkcommon.DefaultConfigAccountAliasesDirName)
}

// SetDefaultLocation set deafault location
func SetDefaultLocation(directory string) {
	sdkcommon.DefaultConfigDirName = directory
	uDir, _ := homedir.Dir()
	tronCTLDir := path.Join(uDir, sdkcommon.DefaultConfigDirName, sdkcommon.DefaultConfigAccountAliasesDirName)
	if _, err := os.Stat(tronCTLDir); os.IsNotExist(err) {
		os.MkdirAll(tronCTLDir, 0700)
	}
}

// UnlockedKeystore return keystore unlocked
func UnlockedKeystore(from, passphrase string) (*keystore.KeyStore, *keystore.Account, error) {
	sender := common.HexToAddress(from)
	ks := FromAddress(from)
	if ks == nil {
		return nil, nil, fmt.Errorf("could not open local keystore for %s", from)
	}
	account, lookupErr := ks.Find(keystore.Account{Address: sender})
	if lookupErr != nil {
		return nil, nil, fmt.Errorf("could not find %s in keystore", from)
	}
	if unlockError := ks.Unlock(account, passphrase); unlockError != nil {
		return nil, nil, errors.Wrap(ErrNoUnlockBadPassphrase, unlockError.Error())
	}
	return ks, &account, nil
}
