package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	User     string
	Password string
	Database int
	Host     string
	Port     string
}

var redisClient *redis.Client

func GetRedisConnection() (*redis.Client, error) {
	if redisClient != nil {
		return redisClient, nil
	}

	db := os.Getenv("REDIS_USER")

	fmt.Println("REDIS_DB: ", db)

	database, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}

	config := RedisConfig{
		User:     os.Getenv("REDIS_USER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: database,
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
	}

	redisClient, err = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       config.Database,
	}), nil

	return redisClient, err
}
