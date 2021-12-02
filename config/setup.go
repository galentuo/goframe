package config

import (
	"log"
	"strings"

	viperLib "github.com/spf13/viper"
)

func Setup(name, configPath, envPrefix, envSeperatorChar string) {
	viperLib.SetConfigName(name)
	viperLib.AddConfigPath(configPath)
	err := viperLib.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}
	viperLib.SetEnvPrefix(envPrefix)

	// Define Replacer
	replacer := strings.NewReplacer(".", envSeperatorChar)
	viperLib.SetEnvKeyReplacer(replacer)
	viperLib.AutomaticEnv()
}
