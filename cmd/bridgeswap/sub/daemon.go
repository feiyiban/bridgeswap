package sub

import (
	"bridgeswap/chains/ethereum"
	"bridgeswap/chains/tron"
	"bridgeswap/config"
	"bridgeswap/logger"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"bridgeswap/controller/core"
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
	chainInfo, err := config.ParseDaemonConfig(cfgFile)
	if err != nil {
		return err
	}
	logger.Debug(fmt.Sprintf("%v", chainInfo))

	sysErr := make(chan error)
	objCore := core.NewCore(sysErr)
	for _, chain := range chainInfo.Chains {
		chainId, errr := strconv.Atoi(chain.ID)
		if errr != nil {
			return errr
		}
		chainConfig := &core.ChainConfig{
			Name:     chain.Name,
			ID:       uint8(chainId),
			Http:     chain.Http,
			Endpoint: chain.Endpoint,
			From:     chain.From,
			Opts:     chain.Opts,
		}

		var newChain core.Chain
		log := logger.Root().New("chain", chainConfig.Name)

		if chain.Type == "ethereum" {
			newChain, err = ethereum.InitializeChain(chainConfig, log, sysErr)
		} else if chain.Type == "tron" {
			newChain, err = tron.InitializeChain(chainConfig, log, sysErr)
		}

		if err != nil {
			logger.Info(err.Error())
			return err
		}

		objCore.AddChain(newChain)
	}

	objCore.Start()

	return nil
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
