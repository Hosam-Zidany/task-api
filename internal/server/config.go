package server

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	JWTSecret string
	JWTExpHrs int
}

var AppConfig Config

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}
	cfg := Config{
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("ENV", "development"),

		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASS", ""),
		DBName: getEnv("DB_NAME", "GoProject2"),

		JWTSecret: getEnv("JWT_SECRET", "change-this"),
	}
	if v := getEnv("JWT_EXP_HOURS", "24"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			cfg.JWTExpHrs = n
		} else {
			cfg.JWTExpHrs = 24
		}
	} else {
		cfg.JWTExpHrs = 24
	}
	AppConfig = cfg
	return cfg
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
