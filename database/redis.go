package database

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

// NewRedisClient initiates Redis client connection
func NewRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}
}
