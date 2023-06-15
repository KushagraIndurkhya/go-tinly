package handlers

import (
	"encoding/json"
	// "fmt"
	"log"
	"net/http"

	"github.com/KushagraIndurkhya/go-tinly/middleware"
	"github.com/KushagraIndurkhya/go-tinly/models"
	"github.com/KushagraIndurkhya/go-tinly/psql"
	redis "github.com/KushagraIndurkhya/go-tinly/redis_layer"
	utils "github.com/KushagraIndurkhya/go-tinly/utills"
	"github.com/gorilla/mux"
)

type request struct {
	Url      string `json:"url"`
	Short    string `json:"short"`
	Comments string `json:"Comments"`
	Medium   string `json:"Medium"`
	Source   string `json:"Source"`
	Campaign string `json:"Campaign"`
	Keyword  string `json:"Keyword"`
	Expiry   int64  `json:"expiry"`
}

func Create(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// w.Header().Set("Access-Control-Allow-Methods", "POST")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")

	id := r.Context().Value(middleware.AuthenticatedUserID).(string)

	// fmt.Printf("ID:%s", id)

	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Print(err)
	}

	// fmt.Print(req.Url)
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
		key := utils.GenKey(5)
		if req.Short != "" {
			key = req.Short
		}
		a := redis.Url_req{Url: req.Url, Created_by: id, Created_at: utils.CurrTime(), Expiry: req.Expiry}
		key, err := redis.Set(&a, key)
		if err != nil {
			log.Print(err)
			resp := make(map[string]interface{})
			resp["status"] = "fail"
			resp["error"] = err.Error()
			json_resp, err := json.Marshal(resp)
			if err != nil {
				log.Print(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(json_resp))

		} else {
			rs := models.URL_INFO_RESPONSE{Url: req.Url, Short: key, Count: 0, Created_by: a.Created_by, Created_at: a.Created_at, Comments: req.Comments, Medium: req.Medium, Source: req.Source, Campaign: req.Campaign, Keyword: req.Keyword}
			resp := make(map[string]interface{})
			log.Print(rs)
			err := psql.Add_URL(a.Created_by, key, rs.Comments, rs.Medium, rs.Source, rs.Campaign, rs.Keyword)
			if err != nil {
				log.Print(err)
			}
			resp["status"] = "success"
			resp["URL_INFO_RESPONSE"] = rs
			json_resp, err := json.Marshal(resp)
			if err != nil {
				log.Print(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(json_resp))
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.AuthenticatedUserID).(string)
	vars := mux.Vars(r)
	key := vars["key"]
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
		err := psql.Del_URL(key)
		if err != nil {
			log.Print(err)
		}
		err = redis.Delete(key)
		if err != nil {
			resp := make(map[string]interface{})
			resp["status"] = "fail"
			resp["error"] = "Something Went Wrong"
			json_resp, err := json.Marshal(resp)
			if err != nil {
				log.Print(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(json_resp))
		} else {
			resp := make(map[string]interface{})
			resp["status"] = "success"
			json_resp, err := json.Marshal(resp)
			if err != nil {
				log.Print(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			w.Write([]byte(json_resp))
		}

	}
}

func CheckAvailability(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.AuthenticatedUserID).(string)
	vars := mux.Vars(r)
	key := vars["key"]
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
		resp := make(map[string]interface{})
		is_available := redis.CheckAvailability(key)
		if is_available {
			resp["status"] = "success"
		}else{
			resp["status"] = "fail"		
		}
		json_resp, err := json.Marshal(resp)
		if err != nil {
			log.Print(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(json_resp))
	}
}

