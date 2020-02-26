package main

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

var db *sql.DB

func main() {
	pgUrl, err := pq.ParseURL("postgres://lvjwvnqa:XhIXr6kW99K-vKH3QrGad-yxlh7ysiIX@rajje.db.elephantsql.com:5432/lvjwvnqa")

	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", pgUrl)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

}
