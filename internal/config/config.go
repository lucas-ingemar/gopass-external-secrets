package config

import (
	"os"
	"strconv"
)

var (
	DEVMODE       = getBoolEnv("DEVMODE")
	API_PORT      = getenv("API_PORT", "3000")
	GIT_COOLDOWN  = getIntEnv("GIT_COOLDOWN", 1)
	GIT_PULL_CRON = getenv("GIT_PULL_CRON", "*/1 * * * *")
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

func getIntEnv(key string, fallback int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	intval, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return intval
}
