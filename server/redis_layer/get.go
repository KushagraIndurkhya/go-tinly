package redis_layer

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/KushagraIndurkhya/go-tinly/models"
	"github.com/go-redis/redis"
)

var (
	counterMutex sync.Mutex
	wg           sync.WaitGroup
	stopFlag     int32
)

func Get(key string) (string, string, error) {

	val, err := Url_db.Get(key).Result()
	if err == redis.Nil {
		log.Printf("%s URL not found", key)
		return "", "", fmt.Errorf("URL NOT FOUND")
	} else if err != nil {
		panic(err)
	}
	// Url_db.Incr(key + "-count").Result()
	// count, err := Url_db.Get(key + "-count").Result()
	// if err == redis.Nil {
	// 	log.Printf("%s count not found", key)
	// } else if err != nil {
	// 	panic(err)
	// }
	// return val, count, nil
	// Push the key into the queue
	_, err = Url_db.LPush("increment-queue", key).Result()
	if err != nil {
		panic(err)
	}

	return val, "", nil
}
func Get_info(info *models.URL_INFO_RESPONSE) (*models.URL_INFO_RESPONSE, error) {
	key := info.Short
	val, err := Url_db.Get(key).Result()
	if err == redis.Nil {
		log.Printf("%s URL not found", key)
		return nil, fmt.Errorf("URL NOT FOUND")
	} else if err != nil {
		log.Print(err)
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
	info.Url = res.Url
	info.Count = count_int
	return info, nil
}
func StartCounterProcessor() {
	wg.Add(1)
	go processIncrementQueue()
}

func StopCounterProcessor() {
	atomic.StoreInt32(&stopFlag, 1)
	wg.Wait()
}

func processIncrementQueue() {
	defer wg.Done()

	for {
		// Check if we should stop processing
		if atomic.LoadInt32(&stopFlag) == 1 {
			return
		}

		// Pop a key from the queue
		key, err := Url_db.RPop("increment-queue").Result()
		if err == redis.Nil {
			// Queue is empty, continue processing
			continue
		} else if err != nil {
			log.Printf("Error while popping from queue: %s", err.Error())
			continue
		}

		// Increment the counter for the key
		counterMutex.Lock()
		_, err = Url_db.Incr(key + "-count").Result()
		counterMutex.Unlock()

		if err != nil {
			log.Printf("Error while incrementing counter for key %s: %s", key, err.Error())
		}
	}
}
