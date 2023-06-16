package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shinobu-uwu/go-challenge-1/server/routes"
)

const SCHEMA = "CREATE TABLE IF NOT EXISTS prices (id INTEGER PRIMARY KEY, price TEXT NOT NULL, timestamp INTEGER NOT NULL);"

func main() {
	db, err := sql.Open("sqlite3", "./prices.db")

	if err != nil {
		panic(err)
	}

	defer db.Close()
	_, err = db.Exec(SCHEMA)

	if err != nil {
		panic(err)
	}

	log.Default().Println("Database initialized")
	http.HandleFunc("/cotacao", routes.Cotacao)
	log.Default().Println("App started")
	http.ListenAndServe(":8080", nil)
}
