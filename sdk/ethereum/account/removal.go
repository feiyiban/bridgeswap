package account

import (
	"fmt"
	"os"
	"path"

	"bridgeswap/sdk/tron/common"
	"bridgeswap/sdk/tron/store"

	"github.com/mitchellh/go-homedir"
)

// RemoveAccount - removes an account from the keystore
func RemoveAccount(name string) error {
	accountExists := store.DoesNamedAccountExist(name)

	if !accountExists {
		return fmt.Errorf("account %s doesn't exist", name)
	}

	uDir, _ := homedir.Dir()
	tronCTLDir := path.Join(uDir, common.DefaultConfigDirName, common.DefaultConfigAccountAliasesDirName)
	accountDir := fmt.Sprintf("%s/%s", tronCTLDir, name)
	os.RemoveAll(accountDir)

	return nil
}
