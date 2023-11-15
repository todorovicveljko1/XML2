package config

import "os"

type Config struct {
	Address     string
	MongoDBURI  string
	NatsAddress string
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
		Address:     getOsValueOrDefault("RESERVATION_ADDRESS", "localhost:8003"),
		MongoDBURI:  getOsValueOrDefault("MONGODB_URI", "mongodb://localhost:27017/"),
		NatsAddress: getOsValueOrDefault("NATS_ADDRESS", "nats://localhost:4222"),
	}
}
