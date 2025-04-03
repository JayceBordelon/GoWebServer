package config

import (
	"encoding/base64"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	UploadPath    string
	Env           string
	EncryptionKey []byte
}

var AppConfig *Config

func LoadConfig() {
	_ = godotenv.Load()

	port := getEnv("PORT", "8080")
	uploadPath := getEnv("UPLOAD_PATH", "./uploads")
	env := getEnv("APP_ENV", "development")
	key := getEnv("ENCRYPTION_KEY", "")
	if key == "" {
		panic("ENCRYPTION_KEY is not set")
	}

	decodedKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		panic("ENCRYPTION_KEY is not valid base64")
	}
	if len(decodedKey) != 32 {
		panic("ENCRYPTION_KEY must decode to exactly 32 bytes (AES-256)")
	}

	AppConfig = &Config{
		Port:          port,
		UploadPath:    uploadPath,
		Env:           env,
		EncryptionKey: decodedKey,
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
