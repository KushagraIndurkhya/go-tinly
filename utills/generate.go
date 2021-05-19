package utills

import (
	"math/rand"
	"time"
)

const set = "ABCDEFGHIJKLMNOPQRSTUVWXTZabcdefghiklmnopqrstuvwxyz0123456789"

func GenKey(size int) string {
	rand.Seed(time.Now().Unix())
	res := ""
	for i := 0; i < 5; i++ {
		res += string(set[rand.Intn(len(set))])
	}
	return res
}
