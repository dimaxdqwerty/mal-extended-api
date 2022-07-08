package db

import (
	"github.com/go-redis/redis"
	"log"
	"mal-extendet-api/models"
)

var client *redis.Client

func GetRedisClient() *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr: models.GetDBHost() + ":" + models.GetDBPort(),
		DB:   1,
	})

	_, err := client.Ping().Result()
	handleErr(err)
	return client
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
