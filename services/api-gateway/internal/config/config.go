package config

import (
	"os"
)

type Config struct {
	Port              string
	AuthServiceURL    string
	ProductServiceURL string
	BiddingServiceURL string
}

func LoadConfig() (*Config, error) {	
	return &Config{
		Port:              os.Getenv("GATEWAY_PORT"),
		AuthServiceURL:    os.Getenv("AUTH_SERVICE_URL"),
		ProductServiceURL: os.Getenv("PRODUCT_SERVICE_URL"),
		BiddingServiceURL: os.Getenv("BIDDING_SERVICE_URL"),
	}, nil 
}
