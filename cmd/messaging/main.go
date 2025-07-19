package main

import (
	"fmt"
	"log"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/config"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/handlers"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	config := config.NewConfig("config", "yaml", "./config/messaging")
	logger, err := middlewares.InitZap()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middlewares.LoggerMiddleware(logger))

	router := r.Group("/api/v1")
	messageHandler := handlers.NewMessageHandler()
	messageHandler.Register(router)

	logger.Info("Server is running", zap.String("host", config.Server.Host), zap.Int("port", config.Server.Port))

	if err := r.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)); err != nil {
		logger.Fatal("Failed to run server", zap.Error(err))
	}
}
