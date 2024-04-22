package config

import (
	"os"
)

const (
	CONFIG_KEY_HOST_NAME   = "RETFIG_HOST_NAME"
	CONFIG_KEY_API_VERSION = "RETFIG_API_VERSION"
	CONFIG_KEY_PORT        = "RETFIG_PORT"
)

var CONFIG *Config = NewConfig()

type Config struct {
	Host       string
	Port       string
	ApiVersion string
}

func NewConfig() *Config {
	host := GetEnv(CONFIG_KEY_HOST_NAME, "localhost")
	port := GetEnv(CONFIG_KEY_PORT, "8080")
	apiVersion := GetEnv(CONFIG_KEY_API_VERSION, "v1")

	return &Config{
		Host:       host,
		Port:       port,
		ApiVersion: apiVersion,
	}
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
