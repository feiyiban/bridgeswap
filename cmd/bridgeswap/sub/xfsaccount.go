/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package sub

import (
	"bridgeswap/sdk/xfsgo/account"
	"bridgeswap/sdk/xfsgo/store"
	"fmt"

	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var (
	xfsAccountCmd = &cobra.Command{
		Use:                   "xfsaccount",
		Short:                 "A brief description of your command",
		DisableFlagsInUseLine: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("account called")
		},
	}
	walletListCommond = &cobra.Command{
		Use:   "list",
		Short: "List all the local accounts",
		RunE:  runWalletList,
	}
	walletGenerateCommond = &cobra.Command{
		Use:                   "generate",
		DisableFlagsInUseLine: true,
		Short:                 "generate xfsgo account address",
		RunE:                  runWalletGenerate,
	}
	walletExportCommand = &cobra.Command{
		Use:                   "export <address>",
		DisableFlagsInUseLine: true,
		Short:                 "export xfsgo account <address>",
		RunE:                  runWalletExport,
	}
	walletImportCommand = &cobra.Command{
		Use:   "import-private-key <secp256k1_PRIVATE_KEY> [ACCOUNT_NAME]",
		Short: "Import an existing keystore key (only accept secp256k1 private keys)",
		Args:  cobra.RangeArgs(1, 2),
		RunE:  runWalletImport,
	}
)

func init() {
	rootCmd.AddCommand(xfsAccountCmd)
	xfsAccountCmd.AddCommand(walletListCommond)
	xfsAccountCmd.AddCommand(walletGenerateCommond)
	xfsAccountCmd.AddCommand(walletImportCommand)
	xfsAccountCmd.AddCommand(walletExportCommand)

}

func runWalletList(cmd *cobra.Command, args []string) error {

	store.DescribeLocalAccounts()
	return nil
}

func runWalletGenerate(cmd *cobra.Command, args []string) error {

	return nil
}

func runWalletExport(cmd *cobra.Command, args []string) error {
	return nil
}

func runWalletImport(cmd *cobra.Command, args []string) error {
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
		fmt.Printf("Xfsgo Address: %s\n", addr)
	}
	return err
}
