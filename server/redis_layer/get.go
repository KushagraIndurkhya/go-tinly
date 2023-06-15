package redis_layer

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/KushagraIndurkhya/go-tinly/models"
	"github.com/go-redis/redis"
)

func Get(key string) (string, string, error) {

	val, err := Url_db.Get(key).Result()
	if err == redis.Nil {
		log.Printf("%s URL not found", key)
		return "", "", fmt.Errorf("URL NOT FOUND")
	} else if err != nil {
		panic(err)
	}
	Url_db.Incr(key + "-count").Result()
	// fmt.Println(result)

	count, err := Url_db.Get(key + "-count").Result()
	if err == redis.Nil {
		log.Printf("%s count not found", key)
	} else if err != nil {
		panic(err)
	}
	return val, count, nil
}
func Get_info(key string) (*models.URL_INFO_RESPONSE, error) {

	val, err := Url_db.Get(key).Result()
	if err == redis.Nil {
		log.Printf("%s URL not found", key)
		return nil, fmt.Errorf("URL NOT FOUND")
	} else if err != nil {
		log.Fatal(err)
	}
	count, err := Url_db.Get(key + "-count").Result()
	if err == redis.Nil {
		log.Printf("%s count not found", key)
	} else if err != nil {
		panic(err)
	}
	var res Url_req
	json.Unmarshal([]byte(val), &res)
	count_int, _ := strconv.Atoi(count)
	rs := models.URL_INFO_RESPONSE{Url: res.Url, Short: key, Count: count_int, Created_by: res.Created_by, Created_at: res.Created_at}
	return &rs, nil
}
