package handlers

import (
	"net/http"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/config"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/middlewares"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/models"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/util"
	"github.com/gin-gonic/gin"
)

type IMessageHandler interface {
	Register(r gin.IRouter)
}

type MessageHandler struct {
	config *config.Config
}

func NewMessageHandler(config *config.Config) IMessageHandler {
	return &MessageHandler{config: config}
}

func (m *MessageHandler) Register(r gin.IRouter) {
	group := r.Group("/message")
	group.Use(middlewares.JWTMiddleware(m.config))

	group.GET("/", func(c *gin.Context) {
		logger := util.LoggerFromContext(c.Request.Context())
		logger.Info("TEST!")
		c.JSON(http.StatusOK, models.NewOkResponse(nil))
	})
}
