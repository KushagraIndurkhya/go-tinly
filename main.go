package main

import (
	"log"
	"net/http"

	"github.com/KushagraIndurkhya/go-tinly/handlers"
	"github.com/KushagraIndurkhya/go-tinly/middleware"
)

func main() {

	app := new(App)
	Initialize(app)

	app.Router.HandleFunc("/r/{Short}", handlers.Redirect)
	app.Router.HandleFunc("/app/api/create", middleware.Auth(handlers.Create))
	app.Router.HandleFunc("/app/login", handlers.HandleLogin)
	app.Router.HandleFunc("/app/callback-gl", handlers.HandleCallback)
	app.Router.HandleFunc("/app/dash", middleware.Auth(handlers.Dash))
	app.Router.HandleFunc("/app", handlers.Test)
	//ToDo:info api "/info/{short}"
	// /api/user
	log.Fatalln(http.ListenAndServe(":8080", app.Router))

}
