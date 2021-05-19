package db

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

type Redis_Client struct {
	Url_db *redis.Client
}

func Connect() (*Redis_Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("DB Connected")
	res := &(Redis_Client{client})
	return res, nil
}
