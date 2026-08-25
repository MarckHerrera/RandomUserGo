package config

import (
	"os"
)

type Config struct {
	Port string 
	EndpointProviderA string
	EndpointProviderB string
	EndpointProviderC string
}

func Load() Config {
	return Config{
		Port: getEnv("PORT", "8080"),
		EndpointProviderA: getEnv("ENDPOINT_PROVIDER_A", "http://localhost:8081"),
		EndpointProviderB: getEnv("ENDPOINT_PROVIDER_B", "http://localhost:8082"),	
		EndpointProviderC: getEnv("ENDPOINT_PROVIDER_C", "http://localhost:8083"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}