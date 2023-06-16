package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Response struct {
	USDBRL Price `json:"USDBRL"`
}

type Price struct {
	ID        int    `json:"-"`
	Price     string `json:"bid"`
	Timestamp int64  `json:"timestamp,string"`
}

func (price *Price) Save() error {
	db, err := sql.Open("sqlite3", "./prices.db")

	if err != nil {
		log.Default().Println("Failed to open database", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	stmt, err := db.PrepareContext(ctx, "INSERT INTO prices (price, timestamp) VALUES (?, ?)")

	if err != nil {
		log.Default().Println("Failed to prepare statement", err)
		return err
	}

	_, err = stmt.Exec(price.Price, price.Timestamp)

	if err != nil {
		log.Default().Println("Failed to insert into database", err)
		return err
	}

	return nil
}
