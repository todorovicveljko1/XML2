package config

import "os"

type Config struct {
	Address    string
	MongoDBURI string
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
		Address:    getOsValueOrDefault("NOTIFICATION_ADDRESS", "localhost:8005"),
		MongoDBURI: getOsValueOrDefault("MONGODB_URI", "mongodb://localhost:27017/"),
	}
}
