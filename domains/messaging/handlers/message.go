package handlers

import (
	"net/http"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/config"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/messaging/usecases"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/middlewares"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type IMessageHandler interface {
	Register(r gin.IRouter)
}

type MessageHandler struct {
	config  *config.Config
	usecase usecases.IWebsocketUseCase
}

func NewMessageHandler(config *config.Config) IMessageHandler {
	return &MessageHandler{
		config:  config,
		usecase: usecases.NewWebsocketUseCase(),
	}
}

func (m *MessageHandler) Register(r gin.IRouter) {
	group := r.Group("/message")
	group.Use(middlewares.JWTMiddleware(m.config))

	group.GET("/ws", m.WebsocketHandler)
}

func (m *MessageHandler) WebsocketHandler(c *gin.Context) {
	ctx := c.Request.Context()
	logger := util.GetLoggerFromContext(ctx)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logger.Error("error while Upgrading websocket connection", zap.Error(err))
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	m.usecase.HandleNewWebsocketClient(ctx, conn)
}
