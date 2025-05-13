package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ENV    string
	Port   string
	ApiKey string

	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	EMailFrom    string
)

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Println("Error loading .env file")
		panic("Error loading .env file")
	}
	ENV = getEnv("ENV", "")
	Port = getEnv("PORT", "")
	ApiKey = getEnv("API_KEY", "")

	SMTPHost = getEnv("SMTP_HOST", "")
	SMTPPort = getEnv("SMTP_PORT", "")
	SMTPUsername = getEnv("SMTP_USERNAME", "")
	SMTPPassword = getEnv("SMTP_PASSWORD", "")
	EMailFrom = getEnv("EMAIL_FROM", "")

}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
