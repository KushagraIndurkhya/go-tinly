package db

import (
	"fmt"

	"github.com/go-redis/redis"
)

func Get(client *Redis_Client, key string) (string, string, error) {

	val, err := client.Url_db.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("no value found")
	} else if err != nil {
		panic(err)
	}
	result, err := client.Url_db.Incr(key + "-count").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	count, err := client.Url_db.Get(key + "-count").Result()
	if err == redis.Nil {
		fmt.Println("no value found")
	} else if err != nil {
		panic(err)
	}
	return val, count, nil

}
