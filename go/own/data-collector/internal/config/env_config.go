package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
    Host     string `envconfig:"DB_HOST"`
    Port     string `envconfig:"DB_PORT"`
    User     string `envconfig:"DB_USER"`
    Password string `envconfig:"DB_PASS"`
}

type OpenAIConfig struct {
    APIKey string `envconfig:"OPENAI_API_KEY"`
}

// All configuration
type Config struct {
    DB       DBConfig    
    OpenAI   OpenAIConfig
    Port     int       `envconfig:"PORT"`
}

// First export environment variables from .env
func InitConfig() Config {
    if err := godotenv.Load(); err != nil {
        fmt.Println("Cannot load .env:", err)
        os.Exit(1)
    }

	// Fill config struct with environment variables
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		fmt.Println("Environment variables mapping error:", err)
		os.Exit(1)
	}
	return cfg
}