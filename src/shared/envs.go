package shared

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvs() map[string]string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return map[string]string{
		"DB_USER":     os.Getenv("DB_USER"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_NAME":     os.Getenv("DB_NAME"),
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_PORT":     os.Getenv("DB_PORT"),
		"DB_SSLMODE":  os.Getenv("DB_SSLMODE"),
		"PORT":        os.Getenv("PORT"),
	}

}
