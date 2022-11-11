// simple api setup
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBConnection struct {
	Conn *pgxpool.Pool
}

func (db *DBConnection) InitConnection() error {
	var err error
	db.Conn, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	return nil
}

func (db *DBConnection) Close() {
	db.Conn.Close()
}

func status(w http.ResponseWriter, r *http.Request) {
	conn := DBConnection{}
	err := conn.InitConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Fprintln(w, "Server is running...")
	// check database connection
	err = conn.Conn.Ping(context.Background())
	if err != nil {
		fmt.Fprintln(w, "Database connection failed: ", err)
		// create database handler
		log.Println("Initializing database handler...")
		err = conn.InitConnection()
		if err != nil {
			log.Println("failed at retry: ", err)
		} else {
			log.Println("Connected to database when retrying")
		}
	} else {
		fmt.Fprintln(w, "Database connection successful")
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong !")
	log.Println("Pong !")
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage!")
}

func main() {
	var wg sync.WaitGroup
	
	wg.Add(1)

	// setup router
	router := mux.NewRouter()

	
	// setup routes
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/ping", ping).Methods("GET")
	router.HandleFunc("/status", status).Methods("GET")
	
	// make channel for graceful shutdown
	c := make(chan os.Signal, 1)

	go func() {
		defer wg.Done()
		log.Println("Starting server...")
		srv := &http.Server{
			Handler: router,
			Addr:    "0.0.0.0:80",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
		c <- os.Interrupt
		log.Println("Server stopped")
	}()
	log.Println("Server started on port " + os.Getenv("PORT"))
	wg.Wait()
	log.Println("Server stopped")
}
