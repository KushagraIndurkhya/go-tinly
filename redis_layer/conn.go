package redis_layer

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var Url_db *redis.Client

func Connect() {

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Redis Not Connected")
	}
	log.Printf("URL DB Connected")
	Url_db = client
}
