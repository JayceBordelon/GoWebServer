package handler

import (
	"errors"
	"net/http"
	"optifile/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service service.FileService
}

func NewHandler(s service.FileService) *Handler {
	return &Handler{Service: s}
}

func (h *Handler) UploadFile(c *gin.Context) {
	err := h.Service.Upload(c)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("file already exists")):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save file"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "file uploaded successfully"})
}

func (h *Handler) DownloadFile(c *gin.Context) {
	data, filename, err := h.Service.Download(c)
	if err != nil {
		switch {
		case errors.Is(err, errors.New("file not found")):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case errors.Is(err, errors.New("file decryption failed")):
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "unexpected download error"})
		}
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Data(http.StatusOK, "application/octet-stream", data)
}

func (h *Handler) DeleteFile(c *gin.Context) {
	err := h.Service.Delete(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "file deleted"})
}

func (h *Handler) GetServerHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Server is healthy",
	})
}
