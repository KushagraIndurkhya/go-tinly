package main

import (
	redis "github.com/KushagraIndurkhya/go-tinly/db"
	r "github.com/go-redis/redis"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *r.Client
}

func Initialize(a *App) {
	redis.Connect()

	a.DB = redis.Url_db
	a.Router = mux.NewRouter()
}
