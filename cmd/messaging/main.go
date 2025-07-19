package main

import (
	"fmt"
	"log"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/config"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/messaging/handlers"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/messaging/models"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/infrastructures"
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

	db, err := infrastructures.NewDB(config)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Group{})
	db.AutoMigrate(&models.UserChat{})

	redis, err := infrastructures.NewRedis(config)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middlewares.LoggerMiddleware(logger))
	r.Use(middlewares.DBMiddleware(db))
	r.Use(middlewares.RedisMiddleware(redis))

	router := r.Group("/api/v1")
	messageHandler := handlers.NewMessageHandler(config)
	messageHandler.Register(router)

	logger.Info("Server is running", zap.String("host", config.Server.Host), zap.Int("port", config.Server.Port))

	if err := r.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)); err != nil {
		logger.Fatal("Failed to run server", zap.Error(err))
	}
}
