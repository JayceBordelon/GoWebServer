package storage

import (
	"errors"
	"io"
	"optifile/internal/config"
	"optifile/internal/utils"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func SaveFile(c *gin.Context) error {
	uploadPath := config.AppConfig.UploadPath
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		return errors.New("failed to create upload directory")
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return errors.New("failed to get file header")
	}

	filePath := filepath.Join(uploadPath, fileHeader.Filename)

	if _, err := os.Stat(filePath); err == nil {
		return errors.New("file already exists")
	}

	src, err := fileHeader.Open()
	if err != nil {
		return errors.New("failed to open file")
	}
	defer src.Close() // This shit is dope

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return errors.New("failed to read file")
	}

	encrypted, err := utils.Encrypt(fileBytes, config.AppConfig.EncryptionKey)
	if err != nil {
		return errors.New("failed to encrypt file")
	}

	if err := os.WriteFile(filePath, encrypted, 0644); err != nil {
		return errors.New("failed to save file")
	}

	return nil
}

func ServeFile(c *gin.Context) ([]byte, string, error) {
	uploadPath := config.AppConfig.UploadPath
	filename := c.Param("filename")
	fullPath := filepath.Join(uploadPath, filename)

	encrypted, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, "", errors.New("failed to read file")
	}

	decrypted, err := utils.Decrypt(encrypted, config.AppConfig.EncryptionKey)
	if err != nil {
		return nil, "", errors.New("failed to decrypt file")
	}

	return decrypted, filename, nil
}



func DeleteFile(c *gin.Context) error {
	uploadPath := config.AppConfig.UploadPath
	filename := c.Param("filename")
	filePath := filepath.Join(uploadPath, filename)

	if err := os.Remove(filePath); err != nil {
		return errors.New("failed to delete file")
	}

	return nil
}

