package utils

import (
	"os"
)

// Env Load environment variable from .env file
func Env(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}
