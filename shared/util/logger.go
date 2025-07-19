package util

import (
	"context"

	"go.uber.org/zap"
)

type LoggerKey string

const LoggerKeyVal LoggerKey = "logger"

func LoggerFromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(LoggerKeyVal).(*zap.Logger)
	if !ok {
		return zap.L()
	}
	return logger
}
