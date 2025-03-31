package handler

import (
	"go-file-server/internal/service"
	"go-file-server/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
    err := service.Upload(c)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "upload successful"})
}


func DownloadFile(c *gin.Context) {
	filename := c.Param("filename")
	logger.Info.Printf("Download requested for: %s", filename)

	service.Download(c)
}

func DeleteFile(c *gin.Context) {
    service.Delete(c)
}
