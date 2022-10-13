package main
// go application to write data to a timescale database from a csv file

import (
	"log"
	"database/sql"
	"sync"

	_ "github.com/lib/pq"
)

type TimescaleDBHandler struct {
	DB *sql.DB
}

var handler *TimescaleDBHandler

var once sync.Once

func NewTimescaleDBHandler() *TimescaleDBHandler {
	once.Do(func() {
		// connect to database
		db, err := sql.Open("postgres", "postgres://postgres:password@localhost:5432/postgres?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		// create handler
		handler = &TimescaleDBHandler{
			DB: db,
		}
	})
	return handler
}

func (tsdb *TimescaleDBHandler) Close() {
	tsdb.DB.Close()
} 