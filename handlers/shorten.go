package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KushagraIndurkhya/go-tinly/middleware"
	"github.com/KushagraIndurkhya/go-tinly/models"
	"github.com/KushagraIndurkhya/go-tinly/psql"
	redis "github.com/KushagraIndurkhya/go-tinly/redis_layer"
	"github.com/KushagraIndurkhya/go-tinly/utills"
)

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
	id := r.Context().Value(middleware.AuthenticatedUserID).(string)
	if id == "" {
		resp := make(map[string]interface{})
		resp["status"] = "fail"
		resp["error"] = "User Not Logged in"
		json_resp, err := json.Marshal(resp)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(json_resp))

		// http.Redirect(w, r, "/app/login", http.StatusTemporaryRedirect)
	} else {

		a := redis.Url_req{Url: req.Url, Created_by: id, Created_at: utills.CurrTime()}
		key, err := redis.Set(&a)
		if err != nil {
			log.Fatal(err)
			resp := make(map[string]interface{})
			resp["status"] = "fail"
			resp["error"] = "Something Went Wrong"
			json_resp, err := json.Marshal(resp)
			if err != nil {
				log.Fatal(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(json_resp))

		} else {
			rs := models.URL_INFO{Url: req.Url, Short: key, Count: 0, Created_by: a.Created_by, Created_at: a.Created_at}
			resp := make(map[string]interface{})

			err := psql.Add_URL(a.Created_by, key)
			if err != nil {
				log.Fatal(err)
			}
			resp["status"] = "success"
			resp["URL_info"] = rs
			json_resp, err := json.Marshal(resp)
			if err != nil {
				log.Fatal(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(json_resp))
		}
	}
}
