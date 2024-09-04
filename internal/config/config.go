package config

import (
	"fmt"
	// "github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort int
}

func LoadConfig() (*Config, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	return nil, fmt.Errorf("error loading .env file: %w", err)
	// }

	// debug enviroment variables reading
	fmt.Println("Loading environment variables...")

	dbPortStr := os.Getenv("DB_PORT")
	fmt.Printf("DB_PORT: %s\n", dbPortStr)  // Debugging line

	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	}

	serverPortStr := os.Getenv("SERVER_PORT")
	fmt.Printf("SERVER_PORT: %s\n", serverPortStr)  // Debugging line

	serverPort, err := strconv.Atoi(serverPortStr)
	if err != nil {
		return nil, fmt.Errorf("invalid SERVER_PORT: %w", err)
	}

	config := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     dbPort,
		
	}

	// dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	// if err != nil {
	// 	return nil, fmt.Errorf("invalid DB_PORT: %w", err)
	// }
	// config.DBPort = dbPort

	// serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	// if err != nil {
	// 	return nil, fmt.Errorf("invalid SERVER_PORT: %w", err)
	// }
	config.ServerPort = serverPort

	return config, nil
}
