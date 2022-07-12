/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package sub

import (
	"bridgeswap/chains/ethereum/crypto"
	"bridgeswap/chains/ethereum/crypto/secp256k1"
	"bridgeswap/chains/ethereum/keystore"
	"bridgeswap/config"
	"bridgeswap/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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
	if len(args) != 1 {
		return cmd.Help()
	}

	logger.Info("Importing key...")
	var err error

	// check if --ed25519 or --sr25519 is set
	keytype := crypto.Secp256k1Type

	_, err = importPrivKey(keytype, "", args[0], nil)

	if err != nil {
		return fmt.Errorf("failed to import key: %w", err)
	}

	return nil
}

func list(cmd *cobra.Command, args []string) error {
	_, err := listKeys("")
	if err != nil {
		return fmt.Errorf("failed to list keys: %w", err)
	}

	return nil
}

//importPrivKey imports a private key into a keypair
func importPrivKey(keytype, datadir, key string, password []byte) (string, error) {
	if password == nil {
		password = keystore.GetPassword("Enter password to encrypt keystore file:")
	}
	keystorepath, err := keystoreDir(datadir)

	if keytype == "" {
		logger.Info("Using default key type", "type", keytype)
		keytype = crypto.Secp256k1Type
	}

	var kp crypto.Keypair

	if keytype == crypto.Secp256k1Type {
		// Hex must not have leading 0x
		if key[0:2] == "0x" {
			kp, err = secp256k1.NewKeypairFromString(key[2:])
		} else {
			kp, err = secp256k1.NewKeypairFromString(key)
		}

		if err != nil {
			return "", fmt.Errorf("could not generate secp256k1 keypair from given string: %w", err)
		}
	} else {
		return "", fmt.Errorf("invalid key type: %s", keytype)
	}

	fp, err := filepath.Abs(keystorepath + "/" + kp.Address() + ".key")
	if err != nil {
		return "", fmt.Errorf("invalid filepath: %w", err)
	}

	file, err := os.OpenFile(filepath.Clean(fp), os.O_EXCL|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return "", fmt.Errorf("Unable to Open File: %w", err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			logger.Error("import private key: could not close keystore file")
		}
	}()

	err = keystore.EncryptAndWriteToFile(file, kp, password)
	if err != nil {
		return "", fmt.Errorf("could not write key to file: %w", err)
	}

	logger.Info("private key imported", "address", kp.Address(), "file", fp)
	return fp, nil

}

// importKey imports a key specified by its filename to datadir/keystore/
// it saves it under the filename "[publickey].key"
// it returns the absolute path of the imported key file
func importKey(filename, datadir string) (string, error) {
	keystorepath, err := keystoreDir(datadir)
	if err != nil {
		return "", fmt.Errorf("could not get keystore directory: %w", err)
	}

	importdata, err := ioutil.ReadFile(filepath.Clean(filename))
	if err != nil {
		return "", fmt.Errorf("could not read import file: %w", err)
	}

	ksjson := new(keystore.EncryptedKeystore)
	err = json.Unmarshal(importdata, ksjson)
	if err != nil {
		return "", fmt.Errorf("could not read file contents: %w", err)
	}

	keystorefile, err := filepath.Abs(keystorepath + "/" + ksjson.Address[2:] + ".key")
	if err != nil {
		return "", fmt.Errorf("could not create keystore file path: %w", err)
	}

	err = ioutil.WriteFile(keystorefile, importdata, 0600)
	if err != nil {
		return "", fmt.Errorf("could not write to keystore directory: %w", err)
	}

	logger.Info("successfully imported key", "address", ksjson.Address, "file", keystorefile)
	return keystorefile, nil
}

// listKeys lists all the keys in the datadir/keystore/ directory and returns them as a list of filepaths
func listKeys(datadir string) ([]string, error) {
	keys, err := getKeyFiles(datadir)
	if err != nil {
		return nil, err
	}

	fmt.Printf("=== Found %d keys ===\n", len(keys))
	for i, key := range keys {
		fmt.Printf("[%d] %s\n", i, key)
	}

	return keys, nil
}

// getKeyFiles returns the filenames of all the keys in the datadir's keystore
func getKeyFiles(datadir string) ([]string, error) {
	keystorepath, err := keystoreDir(datadir)
	if err != nil {
		return nil, fmt.Errorf("could not get keystore directory: %w", err)
	}

	files, err := ioutil.ReadDir(keystorepath)
	if err != nil {
		return nil, fmt.Errorf("could not read keystore dir: %w", err)
	}

	keys := []string{}

	for _, f := range files {
		ext := filepath.Ext(f.Name())
		if ext == ".key" {
			keys = append(keys, f.Name())
		}
	}

	return keys, nil
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
		// uDir, _ := homedir.Dir()
		// // datadir not specified, use default
		// keyPath := filepath.Join(uDir, common.DefaultConfigDirName, common.DefaultConfigAccountAliasesDirName, common.DefaultConfigEthereumAccountDirName)
		// keystorepath, err = filepath.Abs(keyPath)
		// if err != nil {
		// 	return "", fmt.Errorf("could not create keystore file path: %w", err)
		// }
		// datadir not specified, use default
		keyPath = config.DefaultEthKeystorePath

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
