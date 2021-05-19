package handlers

import (
	"encoding/json"
	"net/http"

	redis "github.com/KushagraIndurkhya/go-tinly/db"
	"github.com/gorilla/mux"
)

func Redirect(db *redis.Redis_Client) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		val, _, _ := redis.Get(db, vars["Short"])

		var res redis.Url_req
		json.Unmarshal([]byte(val), &res)
		http.Redirect(w, r, res.Url, http.StatusTemporaryRedirect)
	}
	return http.HandlerFunc(fn)
}
