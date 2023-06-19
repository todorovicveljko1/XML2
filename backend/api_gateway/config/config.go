package config

import "os"

type Config struct {
	Address     string
	AuthAddress string
	AccAddress  string
	ResAddress  string
	NatsAddress string
	RetAddress  string
	Production  bool
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
		AccAddress:  getOsValueOrDefault("ACCOMMODATION_ADDRESS", "localhost:8002"),
		ResAddress:  getOsValueOrDefault("RESERVATION_ADDRESS", "localhost:8003"),
		RetAddress:  getOsValueOrDefault("RATING_ADDRESS", "localhost:8004"),
		NatsAddress: getOsValueOrDefault("NATS_ADDRESS", "localhost:4222"),
		Production:  getOsValueOrDefault("PRODUCTION", "false") == "true",
	}
}
