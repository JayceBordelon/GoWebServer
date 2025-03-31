package storage

import (
	"go-file-server/internal/config"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SaveFile(c *gin.Context) error {
	uploadPath := config.AppConfig.UploadPath

	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return err
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	return c.SaveUploadedFile(file, filepath.Join(uploadPath, file.Filename))
}

func ServeFile(c *gin.Context) {
	uploadPath := config.AppConfig.UploadPath
	filename := c.Param("filename")
	c.File(filepath.Join(uploadPath, filename))
}

func DeleteFile(c *gin.Context) {
	uploadPath := config.AppConfig.UploadPath
	filename := c.Param("filename")
	os.Remove(filepath.Join(uploadPath, filename))
	c.JSON(http.StatusOK, gin.H{"message": "file deleted"})
}
