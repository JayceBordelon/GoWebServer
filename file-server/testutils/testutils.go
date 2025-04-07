package testutils

import (
	"optifile/internal/config"
	"os"
)

func InitTestConfig() {
	os.Setenv("PORT", "9999")
	os.Setenv("UPLOAD_PATH", "./test_uploads")
	os.Setenv("APP_ENV", "test")
	os.Setenv("ENCRYPTION_KEY", "yDRyl9nCL9ME0t9KuzEKM7XYJYRHtEJQjL3RPq88q6s=")
	config.LoadConfig()
}
