package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // or from config
	})
}

func SetUserOnline(userID string) error {
	return Rdb.Set(context.TODO(), "presence:"+userID, "online", 5*time.Minute).Err()
}

func SetUserOffline(userID string) error {
	if err := Rdb.Del(context.TODO(), fmt.Sprintf("presence:%s", userID)).Err(); err != nil {
		return err
	}
	return SetLastSeen(userID) // Save last seen when going offline
}

func GetUserStatus(userID string) (string, error) {
	val, err := Rdb.Get(context.TODO(), "presence:"+userID).Result()
	if err == redis.Nil {
		return "offline", nil
	}
	return val, err
}

func SetLastSeen(userID string) error {
	key := fmt.Sprintf("lastseen:%s", userID)
	timestamp := time.Now().Format(time.RFC3339)
	return Rdb.Set(context.TODO(), key, timestamp, 0).Err()
}

func GetLastSeen(userID string) (string, error) {
	key := fmt.Sprintf("lastseen:%s", userID)
	return Rdb.Get(context.TODO(), key).Result()
}
