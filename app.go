package main

import (
	"github.com/KushagraIndurkhya/go-tinly/config"
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

func Initialize(a *App) {
	config.Init()
	redis.Connect()
	psql.Connect()

	a.Redis_DB = redis.Url_db
	a.PSQL_DB = psql.PSQL_DB
	a.Router = mux.NewRouter()
}
