package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	redis "github.com/KushagraIndurkhya/go-tinly/redis_layer"
	"github.com/KushagraIndurkhya/go-tinly/utills"
)

type response struct {
	Url   string `json:"url"`
	Short string `json:"short_url"`
	// Count      int    `json:"Hits"`
	Created_by string    `json:"Created_by,omitempty"`
	Created_at time.Time `json:"Created_at"`
}
type request struct {
	Url string `json:"url"`
}

func Create(w http.ResponseWriter, r *http.Request) {

	var req request
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		panic(err)
	}
	a := redis.Url_req{Url: req.Url, Count: 0, Created_by: "me"}
	key, _ := redis.Set(&a)

	rs := response{Url: req.Url, Short: key, Created_at: utills.CurrTime()}
	res, err := json.Marshal(&rs)
	if err != nil {
		fmt.Println("error in encoding")
	}
	w.Write(res)
}
