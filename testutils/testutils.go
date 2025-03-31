package testutils

import (
	"go-file-server/internal/config"
	"os"
)

func InitTestConfig() {
	os.Setenv("PORT", "9999")
	os.Setenv("UPLOAD_PATH", "./test_uploads")
	os.Setenv("APP_ENV", "test")
	config.LoadConfig()
}
