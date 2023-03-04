package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Dotenv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading environment")
	}
	return os.Getenv(key)
}
