package psql

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var PSQL_DB *gorm.DB

const (
	DB_DSN = "host=localhost port=5432 user=postgres sslmode=disable dbname=test password=stark "
)

func Connect() {
	db, err := gorm.Open("postgres", DB_DSN)
	if err != nil {
		fmt.Printf("Cannot connect to  database\n")
		log.Fatal("Postgres Connection error:", err)
	} else {
		log.Printf("Users DB Connected")
	}
	PSQL_DB = db
}
