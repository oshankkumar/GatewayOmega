package config

import (
	"github.com/oshankkumar/GatewayOmega/utils"
	"github.com/spf13/viper"
)

var (
	envMap = map[string]string{
		"env":                "ENV",
		"services.nlu.addr":  "NLU",
		"services.auth.addr": "AUTH",
		"log.level":          "LOG_LEVEL",
		"http.port":          "HTTP_PORT",
	}
)

func init() {
	viper.SetConfigFile(utils.GetConfFile())
	viper.ReadInConfig()
	viper.BindEnv(getEnvPairs()...)
}

func getEnvPairs() []string {
	var envpair []string
	for key, val := range envMap {
		envpair = append(envpair, key, val)
	}
	return envpair
}
