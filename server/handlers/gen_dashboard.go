package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/KushagraIndurkhya/go-tinly/middleware"
	"github.com/KushagraIndurkhya/go-tinly/models"
	psql "github.com/KushagraIndurkhya/go-tinly/psql"
	redis "github.com/KushagraIndurkhya/go-tinly/redis_layer"
)

type resp struct {
	UserInfo models.User                `json:"userInfo"`
	Urls     []models.URL_INFO_RESPONSE `json:"urls"`
}

func Dash(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	id := r.Context().Value(middleware.AuthenticatedUserID).(string)
	if id == "" {
		resp := make(map[string]interface{})
		resp["status"] = "fail"
		resp["error"] = "User Not Logged in"
		json_resp, err := json.Marshal(resp)
		if err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(json_resp))

		// http.Redirect(w, r, "/app/login", http.StatusTemporaryRedirect)
	} else {
		info, _ := psql.Get(id)
		url_arr, _ := psql.Get_URLs(id)
		url_inf_arr := []models.URL_INFO_RESPONSE{}
		for i := range url_arr {
			inf, err := redis.Get_info(&url_arr[i])
			if err != nil {
				psql.Del_URL(url_arr[i].Short)
			} else {
				url_inf_arr = append(url_inf_arr, *inf)
			}
		}

		re := resp{UserInfo: *info, Urls: url_inf_arr}
		json_resp, err := json.Marshal(re)

		if err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write(json_resp)

	}
}
