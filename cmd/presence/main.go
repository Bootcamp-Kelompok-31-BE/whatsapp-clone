package main

import (
	"fmt"
	"log"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/config"
	presencehandler "github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/presence/handlers"
	presenceusecase "github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/presence/usecases"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/infrastructures/redis"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	config := config.NewConfig("config", "yaml", "./config/presence")
	logger, err := middlewares.InitZap()
	if err != nil {
		log.Fatal(err)
	}

	// ✅ INIT REDIS
	redis.InitRedis() // assumes localhost:6379

	// ✅ CREATE USECASE & HANDLER
	presenceUC := presenceusecase.NewPresenceUsecase()
	presenceH := presencehandler.NewPresenceHandler(presenceUC)

	// ✅ ROUTES
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middlewares.LoggerMiddleware(logger))

	r.GET("/presence/:id", presenceH.GetStatus)
	r.POST("/presence/:id/online", presenceH.SetOnline)
	r.POST("/presence/:id/offline", presenceH.SetOffline)
	r.GET("/lastseen/:id", presenceH.GetLastSeen)

	// ✅ RUN SERVER
	logger.Info("Server is running",
		zap.String("host", config.Server.Host),
		zap.Int("port", config.Server.Port),
	)

	if err := r.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)); err != nil {
		logger.Fatal("Failed to run server", zap.Error(err))
	}
}
