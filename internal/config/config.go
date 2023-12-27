package config

import (
	"log"
	"os"
)

type Config struct {
	Port string
}

func New() *Config {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	return &Config{
		Port: port,
	}
}
