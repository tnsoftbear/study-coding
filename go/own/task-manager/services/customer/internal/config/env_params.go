package config

import (
	"os"
	"strconv"
)

func GetIntEnv(key string, defaultValue int) int {
	if value := os.Getenv(key); len(value) == 0 {
		return defaultValue
	} else if intValue, err := strconv.Atoi(value); err == nil {
		return intValue
	} else {
		return defaultValue
	}
}

func GetStrEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); len(value) == 0 {
		return defaultValue
	} else {
		return value
	}
}