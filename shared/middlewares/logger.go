package middlewares

import (
	"context"
	"time"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type traceIDKey string

const (
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

		ctx = context.WithValue(ctx, util.LoggerKeyVal, requestLogger)
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
