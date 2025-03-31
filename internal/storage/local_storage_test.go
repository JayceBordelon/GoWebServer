package storage

import (
	"go-file-server/internal/config"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveAndDeleteFile(t *testing.T) {
	os.Setenv("UPLOAD_PATH", "./test_uploads")
	config.LoadConfig()
	defer os.RemoveAll("./test_uploads")
	testFile := filepath.Join(config.AppConfig.UploadPath, "test.txt")

	// Setup: create test dir & dummy file
	err := os.MkdirAll(config.AppConfig.UploadPath, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create test dir: %v", err)
	}

	content := []byte("hello world")
	err = os.WriteFile(testFile, content, 0644)
	if err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	// Simulate delete
	err = os.Remove(testFile)
	if err != nil {
		t.Errorf("Failed to delete file: %v", err)
	}
}
