package main

import (
	"net/http"

	"github.com/KushagraIndurkhya/go-tinly/handlers"
)

func main() {

	app := new(App)
	Initialize(app)

	app.Router.HandleFunc("/{Short}", handlers.Redirect(app.DB))
	app.Router.HandleFunc("/api/create", handlers.Create(app.DB))
	http.ListenAndServe(":8080", app.Router)

}
