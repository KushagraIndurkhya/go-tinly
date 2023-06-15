package redis_layer

import (
	"fmt"
	"time"
)

func Set(req *Url_req,key string) (string, error) {
	val, err := Get_Val(req)
	if err != nil {
		fmt.Print(req, "unable to marshal")
		return "", err
	}
	is_available := CheckAvailability(key)
	if !is_available {
		return "", fmt.Errorf("URL ALREADY EXISTS")
	} 
	err = Url_db.Set(key, val, time.Duration(req.Expiry)*time.Second).Err()
	if err != nil {
		fmt.Print(req, "unable to set URL")
		return "", err
	}
	err = Url_db.Set(key+"-count", 0, 0).Err()
	if err != nil {
		fmt.Print(req, "unable to set Count")
		return "", err
	}
	return key, nil
}

func CheckAvailability(key string) bool {
	_, err := Url_db.Get(key).Result()
	if err != nil {
		return true
	}
	return false
}
