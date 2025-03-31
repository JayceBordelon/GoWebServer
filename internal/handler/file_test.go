package handler

import (
	"go-file-server/pkg/logger"
	"go-file-server/testutils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDownloadFileHandler(t *testing.T) {
	testutils.InitTestConfig()
	logger.Init()

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/files/:filename", DownloadFile)

	req, _ := http.NewRequest("GET", "/files/nonexistent.txt", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound && w.Code != http.StatusOK {
		t.Errorf("Expected status 200 or 404, got %d", w.Code)
	}
}
