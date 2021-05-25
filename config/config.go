package config

import (
	"fmt"
	"os"

	"github.com/KushagraIndurkhya/go-tinly/utills"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig *oauth2.Config
	RandomState       string
	PSQL_DSN          string
)

func Init() {
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:  os.Getenv("REDIRECT_URL"),
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
	RandomState = utills.GenKey(7)

	PSQL_DSN = fmt.Sprintf(
		"host=%s port=%s user=%s sslmode=disable dbname=%s password=%s",
		os.Getenv("PSQL_HOST"),
		os.Getenv("PSQL_PORT"),
		os.Getenv("PSQL_USER"),
		os.Getenv("PSQL_DB"),
		os.Getenv("PSQL_PASS"))

}
