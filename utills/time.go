package utills

import "time"

func CurrTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Kolkata")
	now := time.Now().In(loc)

	return now

}
