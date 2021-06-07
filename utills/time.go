package utills

import (
	"time"
)

func CurrTime() int64 {
	// loc, err := time.LoadLocation("Asia/Kolkata")
	// if err != nil {

	// 	log.Print(err)

	// }
	// now := time.Now().In(loc)

	now := time.Now()
	secs := now.Unix()

	return secs

}
