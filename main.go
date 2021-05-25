package main

import (
	"net/http"

	"github.com/KushagraIndurkhya/go-tinly/handlers"
)

func main() {

	app := new(App)
	Initialize(app)

	app.Router.HandleFunc("/{Short}", handlers.Redirect)
	app.Router.HandleFunc("/api/create", handlers.Create)
	//ToDo:info api "/info/{short}"
	// /api/user
	//
	http.ListenAndServe(":8080", app.Router)

}
