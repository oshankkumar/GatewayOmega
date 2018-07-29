package config

import (
	"github.com/oshankkumar/GatewayOmega/utils"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(utils.GetConfFile())
	viper.ReadInConfig()
	viper.BindEnv()
}
