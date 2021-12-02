package config

import (
	gconf "github.com/galentuo/goframe/config"
	"github.com/spf13/viper"
)

func Simple() *simple {
	gconf.Setup("simple", "./configs/", "simple", "_")
	var sConfig simple

	sConfig.Name = viper.GetString("name")
	sConfig.Env = viper.GetString("env")
	sConfig.Log = logConfig{
		Level: viper.GetString("log.level"),
	}
	sConfig.Server = serverConfig{
		Host: viper.GetString("server.host"),
		Port: viper.GetString("server.port"),
	}

	return &sConfig
}
