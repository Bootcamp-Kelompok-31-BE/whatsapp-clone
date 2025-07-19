package middlewares

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type loggerKey string
type traceIDKey string

const (
	LoggerKey  loggerKey  = "logger"
	TraceIDKey traceIDKey = "trace_id"
)

func InitZap() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	return config.Build()
}

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := uuid.New().String()
		ctx := context.WithValue(c.Request.Context(), TraceIDKey, traceID)

		requestLogger := logger.With(zap.String("trace_id", traceID))

		ctx = context.WithValue(ctx, LoggerKey, requestLogger)
		c.Request = c.Request.WithContext(ctx)

		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		requestLogger.Info("incoming request",
			zap.String("trace_id", traceID),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", status),
			zap.Duration("latency", latency),
			zap.String("client_ip", c.ClientIP()),
		)
	}
}

func FromContext(ctx context.Context) *zap.Logger {
	logger, ok := ctx.Value(LoggerKey).(*zap.Logger)
	if !ok {
		return zap.L()
	}
	return logger
}
