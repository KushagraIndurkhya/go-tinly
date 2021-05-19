package main

import (
	"fmt"

	redis "github.com/KushagraIndurkhya/go-tinly/db"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *redis.Redis_Client
}

func Initialize(a *App) {
	client, err := redis.Connect()
	if err != nil {
		fmt.Println(err)
	}

	a.DB = client
	a.Router = mux.NewRouter()
}
