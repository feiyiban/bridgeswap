/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package sub

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:                   "bridgeswap <command> [<options>]",
	DisableFlagsInUseLine: true,
	SilenceErrors:         true,
	SilenceUsage:          true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func helpTmpl() string {
	return `{{with (or .Long .Short)}}{{. | trimTrailingWhitespaces}}

{{end}}{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`
}

func usageTmpl() string {
	return `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}

Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Options:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] -h,--help" for more information about a command.{{end}}
`
}

func init() {
	mFlags := rootCmd.PersistentFlags()
	mFlags.StringVarP(&cfgFile, "config", "C", "", "Set config file")

	rootCmd.SetHelpTemplate(helpTmpl())
	rootCmd.SetUsageTemplate(usageTmpl())
}
