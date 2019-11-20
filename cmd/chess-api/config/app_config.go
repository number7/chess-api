package config

import "github.com/BurntSushi/toml"

type AppConfig struct {
	ApiName    string       `toml:"api-name"`
	RestServer ServerConfig `toml:"server-config"`
}

type ServerConfig struct {
	Host string `toml:"rest-host"`
	Port int    `toml:"rest-port"`
}

func ParseConfig(configFile string) (*AppConfig, error) {
	var appConfig AppConfig

	if _, err := toml.DecodeFile(configFile, &appConfig); err != nil {
		return nil, err
	}
	return &appConfig, nil
}
