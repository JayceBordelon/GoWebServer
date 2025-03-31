package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	_ = os.Setenv("PORT", "9090")
	_ = os.Setenv("UPLOAD_PATH", "/tmp/uploads")
	_ = os.Setenv("APP_ENV", "test")

	LoadConfig()

	if AppConfig == nil {
		t.Fatal("AppConfig is nil after LoadConfig")
	}

	if AppConfig.Port != "9090" {
		t.Errorf("Expected PORT to be 9090, got %s", AppConfig.Port)
	}

	if AppConfig.UploadPath != "/tmp/uploads" {
		t.Errorf("Expected UPLOAD_PATH to be /tmp/uploads, got %s", AppConfig.UploadPath)
	}

	if AppConfig.Env != "test" {
		t.Errorf("Expected APP_ENV to be test, got %s", AppConfig.Env)
	}
}
