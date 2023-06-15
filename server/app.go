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
	"github.com/joho/godotenv"
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
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	config.Init()
	redis.Connect()
	psql.Connect()

	app.Redis_DB = redis.Url_db
	app.PSQL_DB = psql.PSQL_DB
	app.Router = mux.NewRouter()

	app.Router.HandleFunc("/r/{Short}", handlers.Redirect)
	app.Router.HandleFunc("/api/create", middleware.Auth(handlers.Create))
	app.Router.HandleFunc("/api/delete/{key}", middleware.Auth(handlers.Delete))
	app.Router.HandleFunc("/api/check/{key}", middleware.Auth(handlers.CheckAvailability))
	app.Router.HandleFunc("/login", handlers.HandleLogin)
	app.Router.HandleFunc("/callback-gl", handlers.HandleCallback)
	app.Router.HandleFunc("/api/dash", middleware.Auth(handlers.Dash))

	curr_dir, _ := os.Getwd()
	fmt.Printf("Serving %s...\n", curr_dir+"/build")
	buildHandler := http.FileServer(http.Dir(curr_dir + "/build"))
	app.Router.PathPrefix("/").Handler(buildHandler)
	app.Router.HandleFunc("/home", index).Methods("GET")

}
func (app *App) Run() {
	fmt.Printf("Starting Server On Port%s...\n", config.PORT)
	log.Fatalln(http.ListenAndServe(config.PORT, app.Router))
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./build/Index.html")
}

func main() {

	app := new(App)
	app.Init()
	app.Run()

}
