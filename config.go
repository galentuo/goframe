package goframe

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type configReader struct {
	*viper.Viper
}

func NewConfigReader(name, configPath, envPrefix, envSeperatorChar string) configReader {
	v := viper.New()
	v.SetConfigName(name)
	v.AddConfigPath(configPath)
	err := v.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}
	v.SetEnvPrefix(envPrefix)

	// Define Replacer
	replacer := strings.NewReplacer(".", envSeperatorChar)
	v.SetEnvKeyReplacer(replacer)
	v.AutomaticEnv()

	return configReader{v}
}
