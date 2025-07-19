package config

import (
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		DB     *Database `mapstructure:"db"`
		Server *Server   `mapstructure:"server"`
		JWT    *JWT      `mapstructure:"jwt"`
		Redis  *Redis    `mapstructure:"redis"`
	}

	Database struct {
		Host         string `mapstructure:"host"`
		Port         int    `mapstructure:"port"`
		User         string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		DatabaseName string `mapstructure:"database_name"`
		SslMode      string `mapstructure:"sslmode"`
		TimeZone     string `mapstructure:"timezone"`

		MaxIdleConns             int `mapstructure:"max_idle_conns"`
		MaxOpenConns             int `mapstructure:"max_open_conns"`
		ConnMaxLifetimeInSeconds int `mapstructure:"conn_max_lifetime_in_seconds"`
	}

	Server struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	}

	JWT struct {
		SecretKey string `mapstructure:"secretkey"`
		ExpiredAt int    `mapstructure:"expired_at"`
	}

	Redis struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Password string `mapstructure:"password"`
		DB       int    `mapstructure:"db"`
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
