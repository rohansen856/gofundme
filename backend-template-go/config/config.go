package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type AppConfig struct {
	App struct {
		Name      string
		Port      string
		JWTSecret string
	}
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
		DSN      string
	}
}

var Config AppConfig

func InitConfig(DevMode bool) *AppConfig {
	if DevMode {
		if err := godotenv.Load(".env"); err != nil {
			log.Error().Err(err).Msg("Error loading .env file")
		}
	}

	// Load values from environment variables
	Config.App.Name = getEnv("APP_NAME", "defaultAppName")
	Config.App.Port = getEnv("PORT", "8080")
	Config.App.JWTSecret = getEnv("JWT_SECRET", "defaultJWTSecret")

	Config.DB.Host = getEnv("DB_HOST", "localhost")
	Config.DB.Port = getEnv("DB_PORT", "5432")
	Config.DB.User = getEnv("DB_USER", "your_db_user")
	Config.DB.Password = getEnv("DB_PASSWORD", "your_db_password")
	Config.DB.Name = getEnv("DB_NAME", "your_db_name")
	Config.DB.SSLMode = getEnv("DB_SSL_MODE", "disable")
	Config.DB.DSN = getEnv("DB_DSN", "postgres://your_db_user:your_db_password@localhost:5432/your_db_name?sslmode=disable")

	// Optional: Check if critical database config values are set
	if Config.DB.Host == "" || Config.DB.User == "" || Config.DB.Password == "" {
		log.Fatal().Msg("Missing required environment variables for database connection.")
	}

	return &Config
}

// Helper function to get environment variable or return default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
