package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Cfg *Config
)

type Config struct {
	DBHost      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBPort      string
	DBSsl       string
	DB_TIMEZONE string
	PORT        string
	PRIVATE_KEY string
	PUBLIC_KEY  string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	cfg := &Config{
		DBHost:      os.Getenv("DB_HOST"),
		DBUser:      os.Getenv("DB_USER"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		DBPort:      os.Getenv("DB_PORT"),
		DBSsl:       os.Getenv("DB_SSL"),
		DB_TIMEZONE: os.Getenv("DB_TIMEZONE"),
		PORT:        os.Getenv("PORT"),
		PRIVATE_KEY: os.Getenv("PRIVATE_KEY"),
		PUBLIC_KEY:  os.Getenv("PUBLIC_KEY"),
	}

	return cfg, nil
}
