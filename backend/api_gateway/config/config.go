package config

import "os"

type Config struct {
	Address     string
	AuthAddress string
}

func getOsValueOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConfig() Config {
	return Config{
		Address:     getOsValueOrDefault("API_ADDRESS", "localhost:8000"),
		AuthAddress: getOsValueOrDefault("AUTH_ADDRESS", "localhost:8001"),
	}
}
