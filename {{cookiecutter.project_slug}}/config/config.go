package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Host string
	Port string
}

func LoadConfig() *ServerConfig {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HTTP_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "3000"
	}

	return &ServerConfig{
		Port: port,
		Host: host,
	}
}
