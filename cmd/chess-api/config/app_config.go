package config

import "github.com/BurntSushi/toml"

type AppConfig struct {
	APIName    string       `toml:"api-name"`
	RestServer ServerConfig `toml:"server-config"`
	DBInfo     DBInfo       `toml:"db-config"`
}

type ServerConfig struct {
	Host string `toml:"rest-host"`
	Port int    `toml:"rest-port"`
}

type DBInfo struct {
	Host     string `toml:"db-host"`
	Port     int    `toml:"db-port"`
	DBName   string `toml:"db-name"`
	Username string `toml:"db-username"`
	Password string `toml:"db-password"`
}

func ParseConfig(configFile string) (*AppConfig, error) {
	var appConfig AppConfig

	if _, err := toml.DecodeFile(configFile, &appConfig); err != nil {
		return nil, err
	}
	return &appConfig, nil
}
