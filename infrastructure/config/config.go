package config

import (
	"os"
)

const (
	CONFIG_KEY_HOST_NAME               = "RETFIG_HOST_NAME"
	CONFIG_KEY_API_VERSION             = "RETFIG_API_VERSION"
	CONFIG_KEY_PORT                    = "RETFIG_PORT"
	CONFIG_KEY_JWT_EXPIRES_IN_DURATION = "RETFIG_JWT_EXPIRES_IN_DURATION"
	CONFIG_KEY_JWT_ISSUER              = "RETFIG_JWT_ISSUER"
)

var CONFIG *Config = NewConfig()

type Config struct {
	Host       string
	Port       string
	ApiVersion string
	JWTConfig  JWTConfig
}

type JWTConfig struct {
	ExpiresInDuration string
	Issuer            string
}

func NewJWTConfig() *JWTConfig {
	expiresInDuration := GetEnv(CONFIG_KEY_JWT_EXPIRES_IN_DURATION, "24h")
	issuer := GetEnv(CONFIG_KEY_JWT_ISSUER, "retfig.com")

	return &JWTConfig{
		ExpiresInDuration: expiresInDuration,
		Issuer:            issuer,
	}
}

func NewConfig() *Config {
	host := GetEnv(CONFIG_KEY_HOST_NAME, "127.0.0.1")
	port := GetEnv(CONFIG_KEY_PORT, "8080")
	apiVersion := GetEnv(CONFIG_KEY_API_VERSION, "v1")

	return &Config{
		Host:       host,
		Port:       port,
		ApiVersion: apiVersion,
		JWTConfig:  *NewJWTConfig(),
	}
}

func GetEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
