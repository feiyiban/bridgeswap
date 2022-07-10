package sub

import (
	"bridgeswap/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

var (
	daemonCmd = &cobra.Command{
		Use:                   "daemon [options]",
		DisableFlagsInUseLine: true,
		SilenceUsage:          true,
		Short:                 "Start bridge daemon process",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDaemon()
		},
	}
)

func runDaemon() error {
	chains, err := parseDaemonConfig(cfgFile)
	if err != nil {
		return err
	}

	logger.Info(fmt.Sprintf("%v", chains))

	chanSig := make(chan os.Signal)
	signal.Notify(chanSig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
exit:
	select {
	case sig := <-chanSig:
		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			break exit
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
