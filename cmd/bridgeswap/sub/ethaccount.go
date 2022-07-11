/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package sub

import (
	"bridgeswap/chains/ethereum/crypto"
	"bridgeswap/chains/ethereum/crypto/secp256k1"
	"bridgeswap/chains/ethereum/keystore"
	"bridgeswap/cmd/bridgeswap/sub/common"
	"bridgeswap/logger"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var (
	ethAccountCmd = &cobra.Command{
		Use:                   "ethaccount",
		Short:                 "A brief description of your command",
		DisableFlagsInUseLine: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	handleGenerateCommond = &cobra.Command{
		Use:                   "generate",
		DisableFlagsInUseLine: true,
		Short:                 "generate bridge keystore",
		RunE:                  generatekey,
	}
	handleImportCommond = &cobra.Command{
		Use:                   "import",
		DisableFlagsInUseLine: true,
		Short:                 "import bridge keystore",
		RunE:                  importkey,
	}
	handleListCommond = &cobra.Command{
		Use:                   "list",
		DisableFlagsInUseLine: true,
		Short:                 "list bridge keystore",
		RunE:                  list,
	}
)

func init() {
	rootCmd.AddCommand(ethAccountCmd)
	ethAccountCmd.AddCommand(handleGenerateCommond)
	ethAccountCmd.AddCommand(handleImportCommond)
	ethAccountCmd.AddCommand(handleListCommond)
}

func generatekey(cmd *cobra.Command, args []string) error {
	logger.Info("Generating keypair...")

	keytype := crypto.Secp256k1Type

	_, err := generateKeypair(keytype, "", nil, "")
	if err != nil {
		return fmt.Errorf("failed to generate key: %w", err)
	}
	return nil
}

func importkey(cmd *cobra.Command, args []string) error {
	return nil
}

func list(cmd *cobra.Command, args []string) error {
	return nil
}

// generateKeypair create a new keypair with the corresponding type and saves it to datadir/keystore/[public key].key
// in json format encrypted using the specified password
// it returns the resulting filepath of the new key
func generateKeypair(keytype, datadir string, password []byte, subNetwork string) (string, error) {
	if password == nil {
		password = keystore.GetPassword("Enter password to encrypt keystore file:")
	}

	if keytype == "" {
		logger.Info("Using default key type", "type", keytype)
		keytype = crypto.Secp256k1Type
	}

	var kp crypto.Keypair
	var err error

	if keytype == crypto.Secp256k1Type {
		// generate secp256k1 keys
		kp, err = secp256k1.GenerateKeypair()
		if err != nil {
			return "", fmt.Errorf("could not generate secp256k1 keypair: %w", err)
		}
	} else {
		return "", fmt.Errorf("invalid key type: %s", keytype)
	}

	keystorepath, err := keystoreDir(datadir)
	if err != nil {
		logger.Info("keystorepath:", "keystoreDir", err)
		return "", fmt.Errorf("could not get keystore directory: %w", err)
	}

	fp, err := filepath.Abs(keystorepath + "/" + kp.Address() + ".key")
	if err != nil {
		return "", fmt.Errorf("invalid filepath: %w", err)
	}

	file, err := os.OpenFile(filepath.Clean(fp), os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return "", err
	}

	defer func() {
		err = file.Close()
		if err != nil {
			logger.Error("generate keypair: could not close keystore file")
		}
	}()

	err = keystore.EncryptAndWriteToFile(file, kp, password)
	if err != nil {
		return "", fmt.Errorf("could not write key to file: %w", err)
	}

	logger.Info("key generated", "address", kp.Address(), "type", keytype, "file", fp)
	return fp, nil
}

// keystoreDir returnns the absolute filepath of the keystore directory given a datadir
// by default, it is ./keys/
// otherwise, it is datadir/keys/
func keystoreDir(keyPath string) (keystorepath string, err error) {
	// datadir specified, return datadir/keys as absolute path
	if keyPath != "" {
		keystorepath, err = filepath.Abs(keyPath)
		if err != nil {
			return "", err
		}
	} else {
		uDir, _ := homedir.Dir()
		// datadir not specified, use default
		keyPath := filepath.Join(uDir, common.DefaultConfigDirName, common.DefaultConfigAccountAliasesDirName, common.DefaultConfigEthereumAccountDirName)
		keystorepath, err = filepath.Abs(keyPath)
		if err != nil {
			return "", fmt.Errorf("could not create keystore file path: %w", err)
		}
	}

	// if datadir does not exist, create it
	if _, err = os.Stat(keyPath); os.IsNotExist(err) {
		err = os.Mkdir(keyPath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	// if datadir/keystore does not exist, create it
	if _, err = os.Stat(keystorepath); os.IsNotExist(err) {
		err = os.Mkdir(keystorepath, os.ModePerm)
		if err != nil {
			return "", err
		}
	}

	return keystorepath, nil
}
