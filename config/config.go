package config

import (
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Db     *Database
		Server *Server
	}

	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DbName   string
		SslMode  string
		TimeZone string
	}

	Server struct {
		Port int
		Host string
	}
)

func NewConfig(configName string, configType string, configPath string) *Config {
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
