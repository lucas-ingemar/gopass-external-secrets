package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/rs/zerolog"
)

var (
	DEVMODE       = getBoolEnv("DEVMODE", false)
	AUTH_ACTIVE   = getBoolEnv("AUTH_ACTIVE", true)
	API_PORT      = getenv("API_PORT", "3000")
	GIT_COOLDOWN  = getIntEnv("GIT_COOLDOWN", 5)
	GIT_PULL_CRON = getenv("GIT_PULL_CRON", "*/15 * * * *")
	AUTH_USER     = getenv("AUTH_USER", "")
	AUTH_PASSWORD = getenv("AUTH_PASSWORD", "")
	LOG_LEVEL     = getLogLevel("LOG_LEVEL", zerolog.InfoLevel)
	GOPASS_PREFIX = getenv("GOPASS_PREFIX", "external-secrets")
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getBoolEnv(key string, fallback bool) bool {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return strings.ToLower(value) == "true"
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

func getLogLevel(envVar string, defaultLevel zerolog.Level) zerolog.Level {
	val := strings.ToLower(os.Getenv(envVar))
	switch val {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		return defaultLevel
	}
}
