package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DB_HOST string
var DB_USERNAME string
var DB_PASSWORD string
var DB_SID string
var DB_PORT string
var JWT_SECRET_KEY []byte

func LoadEnvVars() {
	fmt.Println("Loading .env file")
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error Loading .env file: %s", err)
		return
	}

	DB_HOST = os.Getenv("DB_HOST")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_SID = os.Getenv("DB_SID")
	DB_PORT = os.Getenv("DB_PORT")
	JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))
}
