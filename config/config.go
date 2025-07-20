package config

import (
	// "strings"
	"os"
	"log"
	"sync"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Db     *Database
		Server *Server
		JWT    *JWT
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

	JWT struct {
		SecretKey string
		ExpiredAt int
	}
)


var (
	once   sync.Once
	config *Config
)

func NewConfig(configName string, configType string, configPath string) *Config {
	once.Do(func() {
		
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
		viper.AddConfigPath(configPath)
		viper.AutomaticEnv()
		// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

        if dir, err := os.Getwd(); err == nil {
            log.Println("Current working directory:", dir)
        } else {
            log.Panic("Error getting working directory:", err)
        }


		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}

		err := viper.Unmarshal(&config)
		if err != nil {
			panic(err)
		}
		log.Println("Config loaded successfully.")
	})

	return config
}
