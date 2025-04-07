package service

import (
	"optifile/internal/storage"

	"github.com/gin-gonic/gin"
)

type FileService interface {
	Upload(c *gin.Context) error
	Download(c *gin.Context) ([]byte, string, error)
	Delete(c *gin.Context) error
}

type fileServiceImpl struct{}

func NewFileService() FileService {
	return &fileServiceImpl{}
}

func (fs *fileServiceImpl) Upload(c *gin.Context) error {
	return storage.SaveFile(c)
}

func (fs *fileServiceImpl) Download(c *gin.Context) ([]byte, string, error) {
	return storage.ServeFile(c)
}

func (fs *fileServiceImpl) Delete(c *gin.Context) error {
	return storage.DeleteFile(c)
}
