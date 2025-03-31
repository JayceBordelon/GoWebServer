package main

import (
	"go-file-server/internal/config"
	"go-file-server/internal/handler"
	"go-file-server/pkg/logger"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	config.LoadConfig()
	logger.Init()

	r.POST("/upload", handler.UploadFile)
	r.GET("/files/:filename", handler.DownloadFile)
	r.DELETE("/files/:filename", handler.DeleteFile)

	return r
}

func main() {
	r := setupRouter()
	defer logger.Info.Printf("Server started on port: %s", config.AppConfig.Port)

	r.Run(":" + config.AppConfig.Port)
}
