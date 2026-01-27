package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
	MaxConns int
	MaxIdle  int
}

func LoadConfig() (*DBConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	port, err := strconv.Atoi(getEnv("DB_PORT", "5432"))
	if err != nil {
		port = 5432
	}

	maxConns, _ := strconv.Atoi(getEnv("DB_MAX_CONNS", "10"))
	maxIdle, _ := strconv.Atoi(getEnv("DB_MAX_IDLE", "5"))

	return &DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     port,
		User:     getEnv("DB_USER", ""),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", ""),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
		MaxConns: maxConns,
		MaxIdle:  maxIdle,
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}