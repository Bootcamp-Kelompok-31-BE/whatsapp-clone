package middlewares

import (
	"context"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/constant"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ctx = context.WithValue(ctx, constant.DBKey, db)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}
