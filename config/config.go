package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GRPCAddr        string
	HTTPGatewayAddr string
	DB              string
	DBUser          string
	DBPassword      string
	DBPort          string
	DBHost          string
}

func LoadConfig() *Config {
	dir, _ := os.Getwd()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error when trying to load .env file: %v\nDir: %s", err, dir)
	}

	grpcAddr := getEnv("GRPC_ADDR", "0.0.0.0:50051")
	httpGatewayAddr := getEnv("HTTP_GATEWAY_ADDR", "0.0.0.0:8081")
	dbName := getEnv("POSTGRES_DB", "postgres")
	user := getEnv("POSTGRES_USER", "user")
	password := getEnv("POSTGRES_PASSWORD", "postgres")
	port := getEnv("POSTGRES_PORT", "5432")
	host := getEnv("POSTGRES_HOST", "db")

	return &Config{
		GRPCAddr:        grpcAddr,
		HTTPGatewayAddr: httpGatewayAddr,
		DB:              dbName,
		DBUser:          user,
		DBPassword:      password,
		DBPort:          port,
		DBHost:          host,
	}
}

func getEnv(key, base string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	log.Printf("Key %s: value not find. Return base value: %s", key, base)
	return base
}
