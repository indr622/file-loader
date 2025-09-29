package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var FileBasePath string

func InitConfig() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, using system env")
	}

	FileBasePath = os.Getenv("FILE_BASE_PATH")
	if FileBasePath == "" {
		FileBasePath = "./storage"
	}

	// Pastikan folder ada
	if _, err := os.Stat(FileBasePath); os.IsNotExist(err) {
		if err := os.MkdirAll(FileBasePath, os.ModePerm); err != nil {
			log.Fatalf("❌ Failed to create storage dir: %v", err)
		}
	}
}
