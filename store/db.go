package store

import (
	"fmt"

	"github.com/go-redis/redis"
)

var Client *redis.Client

func InitDB() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "database:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Printf(pong)
		panic(err)
	}
	return client
}
