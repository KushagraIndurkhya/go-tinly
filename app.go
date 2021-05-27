package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KushagraIndurkhya/go-tinly/config"
	"github.com/KushagraIndurkhya/go-tinly/handlers"
	"github.com/KushagraIndurkhya/go-tinly/middleware"
	"github.com/KushagraIndurkhya/go-tinly/psql"
	redis "github.com/KushagraIndurkhya/go-tinly/redis_layer"
	r "github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type App struct {
	Router   *mux.Router
	Redis_DB *r.Client
	PSQL_DB  *gorm.DB
}

type Initialize interface {
	Init()
	Build_routes()
	Run()
}

func (app *App) Init() {
	fmt.Print("Initializing App...\n")
	config.Init()
	redis.Connect()
	psql.Connect()

	app.Redis_DB = redis.Url_db
	app.PSQL_DB = psql.PSQL_DB
	app.Router = mux.NewRouter()

	app.Router.HandleFunc("/r/{Short}", handlers.Redirect)
	app.Router.HandleFunc("/api/create", middleware.Auth(handlers.Create))
	app.Router.HandleFunc("/login", handlers.HandleLogin)
	app.Router.HandleFunc("/callback-gl", handlers.HandleCallback)
	app.Router.HandleFunc("/api/dash", middleware.Auth(handlers.Dash))

	buildHandler := http.FileServer(http.Dir("client/build"))
	app.Router.PathPrefix("/").Handler(buildHandler)
	app.Router.HandleFunc("/home", index).Methods("GET")

}
func (app *App) Run() {
	port := os.Getenv("PORT")
	fmt.Printf("Starting Server On Port%s...\n", port)
	log.Fatalln(http.ListenAndServe(port, app.Router))
}
func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "client/build/Index.html")
}
