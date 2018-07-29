package log

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	lvl, err := logrus.ParseLevel(viper.GetString("log.level"))
	if err == nil {
		logrus.SetLevel(lvl)
	}
}
