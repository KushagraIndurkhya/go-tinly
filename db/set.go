package db

import (
	"fmt"

	"github.com/KushagraIndurkhya/go-tinly/utills"
)

func Set(req *Url_req) (string, error) {
	val, err := Get_Val(req)
	if err != nil {
		fmt.Print(req, "unable to marshal")
		return "", err
	}
	key := utills.GenKey(5)
	err = Url_db.Set(key, val, 0).Err()
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
