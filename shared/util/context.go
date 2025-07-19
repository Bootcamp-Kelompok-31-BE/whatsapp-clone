package util

import (
	"context"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/constant"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetUserIDFromContext(ctx context.Context) (uint, bool) {
	userID, ok := ctx.Value(constant.UserIDKey).(uint)
	return userID, ok
}

func GetUserEmailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(constant.UserEmailKey).(string)
	return email, ok
}

func GetDBFromContext(ctx context.Context) *gorm.DB {
	return ctx.Value(constant.DBKey).(*gorm.DB)
}

func GetLoggerFromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(constant.LoggerKey).(*zap.Logger)
	if !ok {
		return zap.L()
	}
	return logger
}

func GetRedisFromContext(ctx context.Context) *redis.Client {
	return ctx.Value(constant.RedisKey).(*redis.Client)
}
