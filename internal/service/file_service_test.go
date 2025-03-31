package service

import (
	"go-file-server/testutils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUploadReturnsErrorForMissingFile(t *testing.T) {
	testutils.InitTestConfig()
	gin.SetMode(gin.TestMode)

	// Create a dummy POST request without a file
	req := httptest.NewRequest("POST", "/upload", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	err := Upload(c)
	if err == nil {
		t.Error("Expected error for missing file, got nil")
	}
}

func TestDownloadAndDelete(t *testing.T) {
	testutils.InitTestConfig()
	gin.SetMode(gin.TestMode)

	filename := "nonexistent.txt"
	req := httptest.NewRequest("GET", "/files/"+filename, nil)
	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = req
	c.Params = []gin.Param{{Key: "filename", Value: filename}}

	Download(c)
	if w.Code != http.StatusOK && w.Code != http.StatusNotFound {
		t.Errorf("Download returned unexpected status: %d", w.Code)
	}

	req = httptest.NewRequest("DELETE", "/files/"+filename, nil)
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request = req
	c.Params = []gin.Param{{Key: "filename", Value: filename}}

	Delete(c)
	if w.Code != http.StatusOK && w.Code != http.StatusNotFound {
		t.Errorf("Delete returned unexpected status: %d", w.Code)
	}
}
