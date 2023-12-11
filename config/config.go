package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GRPCAddr        string
	HTTPGatewayPort string
	DB              string
	DBUser          string
	DBPassword      string
	DBPort          string
}

func LoadConfig() *Config {
	dir, _ := os.Getwd()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error when trying to load .env file: %v\nDir: %s", err, dir)
	}

	grpcAddr := getEnv("GRPC_ADDR", "localhost:50051")
	httpGatewayPort := getEnv("HTTP_GATEWAY_PORT", "8080")
	dbName := getEnv("POSTGRES_DB", "postgres")
	user := getEnv("POSTGRES_USER", "user")
	password := getEnv("POSTGRES_PASSWORD", "postgres")
	port := getEnv("POSTGRES_PORT", "5432")

	return &Config{
		GRPCAddr:        grpcAddr,
		HTTPGatewayPort: httpGatewayPort,
		DB:              dbName,
		DBUser:          user,
		DBPassword:      password,
		DBPort:          port,
	}
}

func getEnv(key, base string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	log.Printf("Key %s: value not find. Return base value: %s", key, base)
	return base
}
