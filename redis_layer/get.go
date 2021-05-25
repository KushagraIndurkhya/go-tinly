package redis_layer

import (
	"fmt"

	"github.com/go-redis/redis"
)

func Get(key string) (string, string, error) {

	val, err := Url_db.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("no value found")
	} else if err != nil {
		panic(err)
	}
	result, err := Url_db.Incr(key + "-count").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	count, err := Url_db.Get(key + "-count").Result()
	if err == redis.Nil {
		fmt.Println("no value found")
	} else if err != nil {
		panic(err)
	}
	return val, count, nil

}
