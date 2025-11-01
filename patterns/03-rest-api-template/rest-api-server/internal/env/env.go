package env

import (
	"os"
	"strconv"
)

// GetString holt Environment Variable mit Fallback
func GetString(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}

// GetInt holt Environment Variable als Integer mit Fallback
func GetInt(key string, fallback int) int {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}

	intVal, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return intVal
}
