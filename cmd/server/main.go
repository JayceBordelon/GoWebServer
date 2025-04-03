package main

import (
	"optifile/internal/config"
	"optifile/internal/handler"
	"optifile/internal/service"
	"optifile/middleware"
	"optifile/pkg/logger"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	config.LoadConfig()
	logger.Init()
	fileService := service.NewFileService()
	h := handler.NewHandler(fileService)

	r.GET("/health", h.GetServerHealth)

	auth := r.Group("/", middleware.AuthMiddleware())
	{
		auth.POST("/files", h.UploadFile)
		auth.GET("/files/:filename", h.DownloadFile)
		auth.DELETE("/files/:filename", h.DeleteFile)
	}

	return r
}

func main() {
	r := SetupRouter()
	defer logger.Info.Printf("Server started on port: %s", config.AppConfig.Port)

	r.Run(":" + config.AppConfig.Port)
}
