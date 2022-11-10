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
)

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is running...")
	// check database connection
	err := conn.DB.Ping(context.Background())
	if err != nil {
		fmt.Fprintln(w, "Database connection failed")
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

func NewDBConnection() {
	return &DBConnection{}
}

func main() {
	var wg sync.WaitGroup
	log.Println("Starting server...")
	wg.Add(1)

	// create database handler
	err = conn.InitConnection()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}

	// init minio
	err = InitMinioClient(config)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Prepared connection to minio")
	}

	// setup router
	router := mux.NewRouter()

	// setup routes
	router.HandleFunc("/ping", ping).Methods("GET")
	router.HandleFunc("/shutdown", shutdown).Methods("GET")
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/list", listObjects).Methods("GET")
	router.HandleFunc("/remaining", listRemaining).Methods("GET")
	router.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		uploadCsvFromMinio(w, r, config)
	}).Methods("GET")
	router.HandleFunc("/generate", generateGeographyPointsRuns).Methods("GET")
	router.HandleFunc("/refresh", refreshListing).Methods("GET")
	router.HandleFunc("/renameMinio", func(w http.ResponseWriter, r *http.Request) {
		renameMinioObjects(w, r, config)
	}).Methods("GET")
	
	// make channel for graceful shutdown
	c := make(chan os.Signal, 1)

	go func() {
		defer wg.Done()
		// setup server
		srv := &http.Server{
			Handler: router,
			Addr:    "0.0.0.0" + ":" + config.PORT,
			// Good practice: enforce timeouts for servers you create!
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		// start server
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
		c <- os.Interrupt
	}()
	log.Println("Server started on port " + os.Getenv("PORT"))
	wg.Wait()
	log.Println("Server stopped")
}
