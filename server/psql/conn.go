package psql

import (
	"fmt"
	"log"

	"github.com/KushagraIndurkhya/go-tinly/config"
	"github.com/KushagraIndurkhya/go-tinly/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var PSQL_DB *gorm.DB

// const (
// 	DB_DSN = "host=localhost port=5432 user=postgres sslmode=disable dbname=test password=stark "
// )

func Connect() {
	db, err := gorm.Open("postgres", config.PSQL_DSN)
	if err != nil {
		fmt.Printf("Cannot connect to  database\n")
		log.Fatal("Postgres Connection error:", err)
	} else {
		log.Printf("Users DB Connected")
	}
	db.Debug().AutoMigrate(models.User{})
	db.Debug().AutoMigrate(models.User_URL{})
	PSQL_DB = db

}
