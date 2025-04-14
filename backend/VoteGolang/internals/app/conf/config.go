package conf

import (
	"log"
	"os"
)

type Config struct {
	JWTSecret string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
}

func LoadConfig() *Config {
	return &Config{
		JWTSecret: getEnv("JWT_SECRET", "defaultsecret"),
		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "3306"),
		DBUser:    getEnv("DB_USER", "root"),
		DBPass:    getEnv("DB_PASS", "Darhani2004"),
		DBName:    getEnv("DB_NAME", "vote_database"),
	}
}

func getEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Warning: Environment variable %s not set, using default value", key)
		return fallback
	}
	return value
}
