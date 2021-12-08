package goframe

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type configReader struct {
	v *viper.Viper
}

func NewConfigReader(fileName, configPath, envPrefix, envSeperatorChar string) *configReader {
	v := viper.New()
	v.SetConfigName(fileName)
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

	return &configReader{v}
}

// Unsafe not recommended to be used.
// It returns an instance of underlying library.
// The signature might be deprecated if & when
// the underlying library is changed.
func (cr *configReader) Unsafe() *viper.Viper {
	return cr.v
}

func (cr *configReader) GetString(key string) string {
	return cr.v.GetString(key)
}

func (cr *configReader) GetBool(key string) bool {
	return cr.v.GetBool(key)
}

func (cr *configReader) GetInt(key string) int {
	return cr.v.GetInt(key)
}

func (cr *configReader) GetStringSlice(key string) []string {
	return cr.v.GetStringSlice(key)
}

func (cr *configReader) GetIntSlice(key string) []int {
	return cr.v.GetIntSlice(key)
}

func (cr *configReader) GetStringMap(key string) map[string]interface{} {
	return cr.v.GetStringMap(key)
}
