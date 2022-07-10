/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package sub

import (
	"fmt"

	"github.com/spf13/cobra"
)

// accountCmd represents the account command
var xfsAccountCmd = &cobra.Command{
	Use:   "xfsaccount",
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

func init() {
	rootCmd.AddCommand(xfsAccountCmd)
}
