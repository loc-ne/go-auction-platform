package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/loc-ne/go-auction/services/api-gateway/config"
)

func main() {
	log.Println("Starting API Gateway...")

	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}
	
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	r := NewRouter(cfg)

	port := cfg.Port
	
	log.Printf("Gateway listening on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
