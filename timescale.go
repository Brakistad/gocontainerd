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

func NewTimescaleDBHandler(host string, port string, username string, password string, db string) *TimescaleDBHandler {
	once.Do(func() {
		// connect to database
		db, err := sql.Open("postgres", "postgres://"+username+":"+password+"@"+host+":"+port+"/"+db+"?sslmode=disable")
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