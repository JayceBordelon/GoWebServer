package service

import (
	"go-file-server/internal/storage"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) error {
    return storage.SaveFile(c)
}


func Download(c *gin.Context) {
    storage.ServeFile(c)
}

func Delete(c *gin.Context) {
    storage.DeleteFile(c)
}
