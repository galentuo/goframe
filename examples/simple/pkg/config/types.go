package config

type logConfig struct {
	Level string
}

type serverConfig struct {
	Host string
	Port string
}

type simple struct {
	Name   string
	Env    string
	Log    logConfig
	Server serverConfig
}
