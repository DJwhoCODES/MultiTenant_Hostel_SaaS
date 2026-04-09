package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv      string
	Port        string
	MongoURI    string
	MongoDB     string
	JWTSecret   string
	JWTExpiry   string
	RabbitMQURL string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system env")
	}

	AppConfig = &Config{
		AppEnv:      getEnv("APP_ENV", "development"),
		Port:        getEnv("PORT", "8080"),
		MongoURI:    getEnv("MONGO_URI", ""),
		MongoDB:     getEnv("MONGO_DB", "hostel_saas"),
		JWTSecret:   getEnv("JWT_SECRET", "secret"),
		JWTExpiry:   getEnv("JWT_EXPIRY", "24h"),
		RabbitMQURL: getEnv("RABBITMQ_URL", ""),
	}
}

func getEnv(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return val
}
