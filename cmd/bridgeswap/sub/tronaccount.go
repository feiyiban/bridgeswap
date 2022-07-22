/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package sub

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"bridgeswap/sdk/tron/account"
	c "bridgeswap/sdk/tron/common"
	"bridgeswap/sdk/tron/ledger"
	"bridgeswap/sdk/tron/mnemonic"
	"bridgeswap/sdk/tron/store"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	useLedgerWallet        bool
	userProvidesPassphrase bool
	addr                   tronAddress
	passphraseFilePath     string
)

// accountCmd represents the account command
var tronAccountCmd = &cobra.Command{
	Use:   "tronaccount",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("account called")
	},
}

const (
	seedPhraseWarning = "**Important** write this seed phrase in a safe place, " +
		"it is the only way to recover your account if you ever forget your password\n\n"
)

var (
	quietImport         bool
	recoverFromMnemonic bool
	passphrase          string
	blsFilePath         string
	blsShardID          uint32
	blsCount            uint32
	ppPrompt            = fmt.Sprintf(
		"prompt for passphrase, otherwise use default passphrase: \"`%s`\"", c.DefaultPassphrase,
	)
)

func keysSub() []*cobra.Command {
	cmdList := &cobra.Command{
		Use:   "list",
		Short: "List all the local accounts",
		RunE: func(cmd *cobra.Command, args []string) error {
			if useLedgerWallet {
				ledger.ProcessAddressCommand()
				return nil
			}
			store.DescribeLocalAccounts()
			return nil
		},
	}

	cmdLocation := &cobra.Command{
		Use:   "location",
		Short: "Show where `tronctl` keeps accounts & their keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(store.DefaultLocation())
			return nil
		},
	}

	cmdAdd := &cobra.Command{
		Use:   "add <ACCOUNT_NAME>",
		Short: "Create a new keystore key",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if store.DoesNamedAccountExist(args[0]) {
				return fmt.Errorf("account %s already exists", args[0])
			}
			passphrase, err := getPassphraseWithConfirm()
			if err != nil {
				return err
			}
			acc := account.Creation{
				Name:       args[0],
				Passphrase: passphrase,
			}

			if err := account.CreateNewLocalAccount(&acc); err != nil {
				return err
			}
			if !recoverFromMnemonic {
				color.Red(seedPhraseWarning)
				fmt.Println(acc.Mnemonic)
			}
			addr, _ := store.AddressFromAccountName(acc.Name)
			fmt.Printf("Tron Address: %s\n", addr)
			return nil
		},
	}

	cmdRemove := &cobra.Command{
		Use:   "remove <ACCOUNT_NAME>",
		Short: "Remove a key from the keystore",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := account.RemoveAccount(args[0]); err != nil {
				return err
			}
			return nil
		},
	}

	cmdMnemonic := &cobra.Command{
		Use:   "mnemonic",
		Short: "Compute the bip39 mnemonic for some input entropy",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(mnemonic.Generate())
			return nil
		},
	}

	cmdRecoverMnemonic := &cobra.Command{
		Use:   "recover-from-mnemonic [ACCOUNT_NAME]",
		Short: "Recover account from mnemonic",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if store.DoesNamedAccountExist(args[0]) {
				return fmt.Errorf("account %s already exists", args[0])
			}
			passphrase, err := getPassphraseWithConfirm()
			if err != nil {
				return err
			}
			acc := account.Creation{
				Name:       args[0],
				Passphrase: passphrase,
			}
			fmt.Println("Enter mnemonic to recover keys from")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			m := scanner.Text()
			if !bip39.IsMnemonicValid(m) {
				return fmt.Errorf("invalid mnemonic given")
			}

			fmt.Println("Enter mnemonic password [optional]")
			scanner.Scan()
			p := scanner.Text()

			acc.Mnemonic = m
			acc.MnemonicPassphrase = p

			if err := account.CreateNewLocalAccount(&acc); err != nil {
				return err
			}
			fmt.Println("Successfully recovered account from mnemonic!")
			addr, _ := store.AddressFromAccountName(acc.Name)
			fmt.Printf("Tron Address: %s\n", addr)
			return nil
		},
	}

	cmdImportKS := &cobra.Command{
		Use:   "import-ks <KEYSTORE_FILE_PATH> [ACCOUNT_NAME]",
		Args:  cobra.RangeArgs(1, 2),
		Short: "Import an existing keystore key",
		RunE: func(cmd *cobra.Command, args []string) error {
			userName := ""
			if len(args) == 2 {
				userName = args[1]
			}
			passphrase, err := getPassphrase()
			if err != nil {
				return err
			}
			name, err := account.ImportKeyStore(args[0], userName, passphrase)
			if !quietImport && err == nil {
				fmt.Printf("Imported keystore given account alias of `%s`\n", name)
				addr, _ := store.AddressFromAccountName(name)
				fmt.Printf("Tron Address: %s\n", addr)
			}
			return err
		},
	}
	cmdImportKS.Flags().BoolVar(&quietImport, "quiet", false, "do not print out imported account name")

	cmdImportPK := &cobra.Command{
		Use:   "import-private-key <secp256k1_PRIVATE_KEY> [ACCOUNT_NAME]",
		Short: "Import an existing keystore key (only accept secp256k1 private keys)",
		Args:  cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			userName := ""
			if len(args) == 2 {
				userName = args[1]
			}
			passphrase, err := getPassphrase()
			if err != nil {
				return err
			}
			name, err := account.ImportFromPrivateKey(args[0], userName, passphrase)
			if !quietImport && err == nil {
				fmt.Printf("Imported keystore given account alias of `%s`\n", name)
				addr, _ := store.AddressFromAccountName(name)
				fmt.Printf("Tron Address: %s\n", addr)
			}
			return err
		},
	}
	cmdImportPK.Flags().BoolVar(&quietImport, "quiet", false, "do not print out imported account name")

	cmdExportPK := &cobra.Command{
		Use:     "export-private-key <ACCOUNT_ADDRESS>",
		Short:   "Export the secp256k1 private key",
		Args:    cobra.ExactArgs(1),
		PreRunE: validateAddress,
		RunE: func(cmd *cobra.Command, args []string) error {
			passphrase, err := getPassphrase()
			if err != nil {
				return err
			}
			return account.ExportPrivateKey(addr.address, passphrase)
		},
	}

	cmdExportKS := &cobra.Command{
		Use:     "export-ks <ACCOUNT_ADDRESS> <OUTPUT_DIRECTORY>",
		Short:   "Export the keystore file contents",
		Args:    cobra.ExactArgs(2),
		PreRunE: validateAddress,
		RunE: func(cmd *cobra.Command, args []string) error {
			passphrase, err := getPassphrase()
			if err != nil {
				return err
			}
			file, e := account.ExportKeystore(addr.address, args[1], passphrase)
			if file != "" {
				fmt.Println("Exported keystore to", file)
			}
			return e
		},
	}

	return []*cobra.Command{cmdList, cmdLocation, cmdAdd, cmdRemove, cmdMnemonic, cmdRecoverMnemonic, cmdImportKS, cmdImportPK,
		cmdExportKS, cmdExportPK}
}

func init() {
	tronAccountCmd.AddCommand(keysSub()...)
	rootCmd.AddCommand(tronAccountCmd)
}

func validateAddress(cmd *cobra.Command, args []string) error {
	// Check if input valid one address
	var err error
	addr, err = findAddress(args[0])
	return err
}

func findAddress(value string) (tronAddress, error) {
	// Check if input valid one address
	address := tronAddress{}
	if err := address.Set(value); err != nil {
		// Check if input is valid account name
		if acc, err := store.AddressFromAccountName(value); err == nil {
			return tronAddress{acc}, nil
		}
		return address, fmt.Errorf("Invalid address/Invalid account name: %s", value)
	}
	return address, nil
}

// getPassphrase fetches the correct passphrase depending on if a file is available to
// read from or if the user wants to enter in their own passphrase. Otherwise, just use
// the default passphrase. No confirmation of passphrase
func getPassphrase() (string, error) {
	if passphraseFilePath != "" {
		if _, err := os.Stat(passphraseFilePath); os.IsNotExist(err) {
			return "", fmt.Errorf("passphrase file not found at `%s`", passphraseFilePath)
		}
		dat, err := ioutil.ReadFile(passphraseFilePath)
		if err != nil {
			return "", err
		}
		pw := strings.TrimSuffix(string(dat), "\n")
		return pw, nil
	} else if userProvidesPassphrase {
		fmt.Println("Enter passphrase:")
		pass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return "", err
		}
		return string(pass), nil
	} else {
		return c.DefaultPassphrase, nil
	}
}

// getPassphrase fetches the correct passphrase depending on if a file is available to
// read from or if the user wants to enter in their own passphrase. Otherwise, just use
// the default passphrase. Passphrase requires a confirmation
func getPassphraseWithConfirm() (string, error) {
	if passphraseFilePath != "" {
		if _, err := os.Stat(passphraseFilePath); os.IsNotExist(err) {
			return "", fmt.Errorf("passphrase file not found at `%s`", passphraseFilePath)
		}
		dat, err := ioutil.ReadFile(passphraseFilePath)
		if err != nil {
			return "", err
		}
		pw := strings.TrimSuffix(string(dat), "\n")
		return pw, nil
	} else if userProvidesPassphrase {
		fmt.Println("Enter passphrase:")
		pass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return "", err
		}
		fmt.Println("Repeat the passphrase:")
		repeatPass, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return "", err
		}
		if string(repeatPass) != string(pass) {
			return "", errors.New("passphrase does not match")
		}
		fmt.Println("") // provide feedback when passphrase is entered.
		return string(repeatPass), nil
	} else {
		return c.DefaultPassphrase, nil
	}
}
