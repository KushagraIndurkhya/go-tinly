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
	UserInfo models.User       `json:"userInfo"`
	Urls     []models.URL_INFO `json:"urls"`
}

func Dash(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.AuthenticatedUserID).(string)

	if id == "" {
		// ctx := context.WithValue(r.Context(), "redirect", r.URL)
		http.Redirect(w, r, "/app/login", http.StatusTemporaryRedirect)
	} else {
		info, _ := psql.Get(id)
		url_arr, _ := psql.Get_URLs(id)
		url_inf_arr := []models.URL_INFO{}
		for i := range url_arr {

			inf, err := redis.Get_info(url_arr[i])
			if err != nil {
				psql.Del_URL(url_arr[i])
			} else {
				url_inf_arr = append(url_inf_arr, *inf)
			}
		}

		re := resp{UserInfo: *info, Urls: url_inf_arr}
		json_resp, err := json.Marshal(re)

		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write(json_resp)

	}
}
