package handlers

import (
	"net/http"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/models"
	"github.com/gin-gonic/gin"
)

type IMessageHandler interface {
	Register(r gin.IRouter)
}

type MessageHandler struct {
}

func NewMessageHandler() IMessageHandler {
	return &MessageHandler{}
}

func (m *MessageHandler) Register(r gin.IRouter) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.NewOkResponse(nil))
	})
}
