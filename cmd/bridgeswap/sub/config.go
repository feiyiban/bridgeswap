package sub

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	defaultConfigFile = "./config.yaml"
)

type ConfigChains struct {
	Chains []Chain `mapstructure:"chains" json:"chains" yaml:"chains"`
}

type Chain struct {
	Name     string `mapstructure:"name" json:"name" yaml:"name"`             // 链名称
	Type     string `mapstructure:"type" json:"type" yaml:"type"`             // 链类型
	ID       string `mapstructure:"id" json:"id" yaml:"id"`                   // 桥接链ID
	Endpoint string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"` // 对应桥的请求地址
	From     string `mapstructure:"from" json:"from" yaml:"from"`             // 签名地址
	Opts     Opts   `mapstructure:"opts" json:"opts" yaml:"opts"`             // 合约
}

type Opts struct {
	Bridge string `mapstructure:"bridge" json:"bridge" yaml:"bridge"` // 桥接合约
	ERC20  string `mapstructure:"erc20" json:"erc20" yaml:"erc20"`    // 桥接合约
}

func parseDaemonConfig(configFilePath string) (ConfigChains, error) {
	if configFilePath == "" {
		configFilePath = defaultConfigFile
	}
	vip := viper.New()

	vip.SetConfigFile(configFilePath)
	vip.SetConfigType("yaml")

	err := vip.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	chains := ConfigChains{}

	vip.WatchConfig()

	vip.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = vip.Unmarshal(&chains); err != nil {
			fmt.Println(err)
		}
	})

	if err = vip.Unmarshal(&chains); err != nil {
		fmt.Println(err)
		return chains, nil
	}

	return chains, nil
}