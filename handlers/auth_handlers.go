package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/KushagraIndurkhya/go-tinly/config"
	"github.com/KushagraIndurkhya/go-tinly/middleware"
	"github.com/KushagraIndurkhya/go-tinly/models"
	psql "github.com/KushagraIndurkhya/go-tinly/psql"
	"github.com/KushagraIndurkhya/go-tinly/utills"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	url := config.GoogleOauthConfig.AuthCodeURL(config.RandomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("state") != config.RandomState {
		fmt.Print("Invalid State")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := config.GoogleOauthConfig.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		fmt.Print("Could not retrieve token")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Print("Could not create get request")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	//Save user info to db
	content, _ := ioutil.ReadAll(resp.Body)
	s := new(models.UserInfo)
	json.Unmarshal(content, &s)
	psql.Save(s)

	//Set Cookie Token
	cookie, _ := utills.Make_Cookie(s.Id)
	http.SetCookie(w, cookie)
	//Redirect to Home
	http.Redirect(w, r, "/app/dash", http.StatusTemporaryRedirect)
}

func Test(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(`<h1> Go-Tinly </h1>
	<a href="\login"> Google Login</a>`))

}

func Dash(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middleware.AuthenticatedUserID).(string)

	if id == "" {
		// ctx := context.WithValue(r.Context(), "redirect", r.URL)
		http.Redirect(w, r, "/app/login", http.StatusTemporaryRedirect)
	} else {
		UserInfo, _ := psql.Get(id)
		w.Write([]byte(UserInfo.Email))
		fmt.Print(UserInfo)
	}
}
