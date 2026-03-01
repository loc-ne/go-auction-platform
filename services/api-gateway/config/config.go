package config

import (
	"os"
)

type Config struct {
	Port           string
	AuthServiceURL string
}

func LoadConfig() (*Config, error) {	
	return &Config{
		Port:           os.Getenv("GATEWAY_PORT"),
		AuthServiceURL: os.Getenv("AUTH_SERVICE_URL"),
	}, nil 
}
