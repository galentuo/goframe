package goframe

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config is used to read the configs from config files or env variables.
/*
	config := NewConfig(fileName, configPath, envPrefix, envSeparatorChar)
*/
type Config interface {
	// GetString gets a string config
	GetString(key string) string
	// GetBool gets a boolean config
	GetBool(key string) bool
	// GetInt gets an integer config
	GetInt(key string) int
	// GetStringSlice gets a string slice ([]string) config
	GetStringSlice(key string) []string
	// GetIntSlice gets an integer slice ([]int) config
	GetIntSlice(key string) []int
	// GetStringMap gets a string Map (map[string]interface{}) config
	GetStringMap(key string) map[string]interface{}
}

type configReader struct {
	v *viper.Viper
}

// NewConfig is used to create a new instance of Config
// Config reads the configs from config files or env variables.
/*
	config := NewConfig(fileName, configPath, envPrefix, envSeparatorChar)
*/
func NewConfig(fileName, configPath, envPrefix, envSeparatorChar string) *configReader {
	v := viper.New()
	v.SetConfigName(fileName)
	v.AddConfigPath(configPath)
	err := v.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}
	v.SetEnvPrefix(envPrefix)

	// Define Replacer
	replacer := strings.NewReplacer(".", envSeparatorChar)
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

// GetString gets a string config
func (cr *configReader) GetString(key string) string {
	return cr.v.GetString(key)
}

// GetBool gets a boolean config
func (cr *configReader) GetBool(key string) bool {
	return cr.v.GetBool(key)
}

// GetInt gets an integer config
func (cr *configReader) GetInt(key string) int {
	return cr.v.GetInt(key)
}

// GetStringSlice gets a string slice ([]string) config
func (cr *configReader) GetStringSlice(key string) []string {
	return cr.v.GetStringSlice(key)
}

// GetIntSlice gets an integer slice ([]int) config
func (cr *configReader) GetIntSlice(key string) []int {
	return cr.v.GetIntSlice(key)
}

// GetStringMap gets a string Map (map[string]interface{}) config
func (cr *configReader) GetStringMap(key string) map[string]interface{} {
	return cr.v.GetStringMap(key)
}
