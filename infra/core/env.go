package core

import (
	"fmt"
	"os"
)

func GetEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func GetEnvAsInt(key string, defaultVal int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		var value int
		_, err := fmt.Sscanf(valueStr, "%d", &value)
		if err == nil {
			return value
		}
	}
	return defaultVal
}
