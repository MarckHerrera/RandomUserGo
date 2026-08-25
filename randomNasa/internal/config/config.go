package config

import (
	"os"
)

type Config struct {
	Port string
	EndpointA string
}

func Load() Config{
	return Config{
		Port: getEnv("PORT", "8080"),
		EndpointA: getEnv("Endpoint_A", "http://localhost:8081"),
	}
}

func getEnv(key, defaultV string) string{
	if valor, ok := os.LookupEnv(key); ok {
		return valor
	}
	return defaultV
}