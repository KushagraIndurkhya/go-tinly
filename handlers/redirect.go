package handlers

import (
	"encoding/json"
	"net/http"

	redis "github.com/KushagraIndurkhya/go-tinly/redis_layer"
	"github.com/gorilla/mux"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	val, _, _ := redis.Get(vars["Short"])

	var res redis.Url_req
	json.Unmarshal([]byte(val), &res)
	http.Redirect(w, r, res.Url, http.StatusTemporaryRedirect)
}
