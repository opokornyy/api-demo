package config

import (
	"fmt"
	"os"
)

const (
	DATABASE_URL = "DATABASE_URL"
	SERVER_URL   = "SERVER_URL"
)

type Config struct {
	DatabaseURL string
	ServerURL   string
}

func NewConfig() (*Config, error) {
	databaseURL, err := getEnv(DATABASE_URL)
	if err != nil {
		return nil, err
	}

	serverUrl, err := getEnv(SERVER_URL)
	if err != nil {
		return nil, err
	}

	return &Config{
		DatabaseURL: databaseURL,
		ServerURL:   serverUrl,
	}, nil
}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("%s is required", key)
	}
	return value, nil
}
