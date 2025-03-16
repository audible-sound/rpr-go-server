package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// Load Environment Variables
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File!")
	}
}
