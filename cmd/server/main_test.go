package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"optifile/internal/config"
	"optifile/testutils"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter_FileLifecycle(t *testing.T) {
	t.Log("Initializing config and setting up router")
	testutils.InitTestConfig()
	router := SetupRouter()

	filename := "testfile.txt"
	content := "This is a test file."
	uploadPath := config.AppConfig.UploadPath
	defer os.RemoveAll(uploadPath)

	
	t.Run("POST /files - Upload", func(t *testing.T) {
		var buf bytes.Buffer
		writer := multipart.NewWriter(&buf)
		part, err := writer.CreateFormFile("file", filename)
		assert.NoError(t, err)
		_, _ = io.Copy(part, strings.NewReader(content))
		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/files", &buf)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		req.Header.Set("Authorization", "Bearer supersecrettoken")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		t.Logf("Upload response: %d", resp.Code)
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	
	t.Run("GET /files/:filename - Download", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/files/"+filename, nil)
		req.Header.Set("Authorization", "Bearer supersecrettoken")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		t.Logf("Download response: %d", resp.Code)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Equal(t, content, resp.Body.String())
	})

	
	t.Run("DELETE /files/:filename - Delete", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/files/"+filename, nil)
		req.Header.Set("Authorization", "Bearer supersecrettoken")
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		t.Logf("Delete response: %d", resp.Code)
		assert.Equal(t, http.StatusOK, resp.Code)

		_, err := os.Stat(filepath.Join(uploadPath, filename))
		assert.True(t, os.IsNotExist(err), "file should be deleted")
	})

	
	t.Run("GET /health - Health Check", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/health", nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		t.Logf("❤️ Health check response: %d", resp.Code)
		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, resp.Body.String(), "Server is healthy")
	})
}
