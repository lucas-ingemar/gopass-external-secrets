package config

import "os"

var (
	DEVMODE  = getBoolEnv("DEVMODE")
	API_PORT = getenv("API_PORT", "3000")
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getBoolEnv(key string) bool {
	value := os.Getenv(key)
	return len(value) != 0
}
