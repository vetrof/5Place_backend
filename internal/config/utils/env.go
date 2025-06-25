package utils

import (
	"os"
	"strconv"
)

// getEnvOrDefault возвращает значение переменной окружения или значение по умолчанию
func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvIntOrDefault возвращает int значение переменной окружения или значение по умолчанию
func GetEnvIntOrDefault(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
