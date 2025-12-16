package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
)

// RedisConfig represents the configuration for a Redis connection
type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// LoadRedisConfig loads the Redis configuration from a JSON file
func LoadRedisConfig(filePath string) (*RedisConfig, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config RedisConfig
	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// NewRedisClient creates a new Redis client based on the provided configuration
func NewRedisClient(config *RedisConfig) (*redis.Client, error) {
	if config == nil {
		return nil, errors.New("redis config is nil")
	}

	address := config.Host + ":" + strconv.Itoa(config.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}

// GetRedisPassword returns the Redis password from the environment variables
func GetRedisPassword() string {
	password := os.Getenv("REDIS_PASSWORD")
	if password == "" {
		log.Fatal("REDIS_PASSWORD environment variable is not set")
	}
	return password
}

// GetRedisHost returns the Redis host from the environment variables
func GetRedisHost() string {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		log.Fatal("REDIS_HOST environment variable is not set")
	}
	return host
}

// GetRedisPort returns the Redis port from the environment variables
func GetRedisPort() int {
	portStr := os.Getenv("REDIS_PORT")
	if portStr == "" {
		log.Fatal("REDIS_PORT environment variable is not set")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal("REDIS_PORT environment variable is not a valid integer")
	}

	return port
}

// GetRedisDB returns the Redis database from the environment variables
func GetRedisDB() int {
	dbStr := os.Getenv("REDIS_DB")
	if dbStr == "" {
		log.Fatal("REDIS_DB environment variable is not set")
	}

	db, err := strconv.Atoi(dbStr)
	if err != nil {
		log.Fatal("REDIS_DB environment variable is not a valid integer")
	}

	return db
}