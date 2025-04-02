package main

import (
	"go-file-server/testutils"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMainRouterRoutes(t *testing.T) {
	testutils.InitTestConfig()

	_ = os.MkdirAll("./test_uploads", os.ModePerm)
	_ = os.WriteFile("./test_uploads/test.txt", []byte("hello test"), 0644)
	defer os.RemoveAll("./test_uploads")

	gin.SetMode(gin.TestMode)
	router := setupRouter()

	tests := []struct {
		method string
		path   string
	}{
		{"GET", "/health"},
		{"POST", "/upload"},
		{"GET", "/files/test.txt"},
		{"DELETE", "/files/test.txt"},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest(tt.method, tt.path, nil)
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)

		if resp.Code == http.StatusNotFound {
			t.Errorf("Route %s %s is not registered (got 404)", tt.method, tt.path)
		}
	}
}
