package infrastructures

import (
	"context"
	"fmt"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/config"
	"github.com/redis/go-redis/v9"
)

func NewRedis(config *config.Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
