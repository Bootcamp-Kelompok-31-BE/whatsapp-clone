package main

import (
	"fmt"
	"log"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/config"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/middlewares"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/entities"
	mediaHttp "github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/handlers/http"
	mediaRepo "github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/repositories"
	mediaMc "github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/media/usecases"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/infrastructures"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	config := config.NewConfig("config", "yaml", "./config/media")
	db := infrastructures.NewPostgresDatabase(config)
	PostgresDatabase := db
	MediaDatabaseRepo := mediaRepo.NewDatabaseMediaRepository(PostgresDatabase)
	MediaMc           := mediaMc.NewMediaUseCase(MediaDatabaseRepo)
	MediaHttp         := mediaHttp.NewMediaHttp(MediaMc)

	migrateDB := db.GetInstance()
	migrateDB.AutoMigrate(&entities.Media{})
	
	logger, err := middlewares.InitZap()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middlewares.LoggerMiddleware(logger))

	r.GET("/home", MediaHttp.Home)

	media := r.Group("/api/v1/")
	{
		media.GET("/media/find/:name", MediaHttp.GetFile)
		media.POST("/media/upload", MediaHttp.UploadFile)
	}

	logger.Info("Server is running", zap.String("host", config.Server.Host), zap.Int("port", config.Server.Port))

	if err := r.Run(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)); err != nil {
		logger.Fatal("Failed to run server", zap.Error(err))
	}


}
