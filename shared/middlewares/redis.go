package middlewares

import (
	"context"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/constant"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RedisMiddleware(redis *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, constant.RedisKey, redis)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
