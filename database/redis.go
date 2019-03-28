package database

import (
	"log"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func NewRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal(err.Error())
	}
}
