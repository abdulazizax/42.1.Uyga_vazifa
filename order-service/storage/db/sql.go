package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	drName     = "postgres"
	dbUrl      = "localhost"
	dbPort     = 5432
	dbName     = "n9"
	dbUser     = "postgres"
	dbPassword = "abdulaziz1221"
)

func DBConnect() *sqlx.DB {
	dbStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbUrl, dbPort, dbUser, dbPassword, dbName)
	db, err := sqlx.Open(drName, dbStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

