/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package sub

import (
	"bridgeswap/sdk/ethereum/account"
	"bridgeswap/sdk/ethereum/store"
	"fmt"

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
		Short:                 "generate eth account",
		RunE:                  generatekey,
	}
	handleImportCommond = &cobra.Command{
		Use:                   "import",
		DisableFlagsInUseLine: true,
		Short:                 "import eth account key",
		RunE:                  importkey,
	}
	handleListCommond = &cobra.Command{
		Use:                   "list",
		DisableFlagsInUseLine: true,
		Short:                 "list eth account key",
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

	return nil
}

func importkey(cmd *cobra.Command, args []string) error {

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
		fmt.Printf("Ethereum Address: %s\n", addr)
	}
	return err
}

func list(cmd *cobra.Command, args []string) error {

	return nil
}
