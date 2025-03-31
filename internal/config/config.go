package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	UploadPath string
	Env        string
}

var AppConfig *Config

func LoadConfig() {
	_ = godotenv.Load()

	port := getEnv("PORT", "8080")
	uploadPath := getEnv("UPLOAD_PATH", "./uploads")
	env := getEnv("APP_ENV", "development")

	AppConfig = &Config{
		Port:       port,
		UploadPath: uploadPath,
		Env:        env,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
