package config

import "os"

type Config struct {
	Address     string
	MongoDBURI  string
	Secret      string
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
		Address:     getOsValueOrDefault("AUTH_ADDRESS", "localhost:8001"),
		MongoDBURI:  getOsValueOrDefault("MONGODB_URI", "mongodb://localhost:27017/"),
		Secret:      getOsValueOrDefault("SECRET", "secret"),
		NatsAddress: getOsValueOrDefault("NATS_ADDRESS", "nats://localhost:4222"),
	}
}
