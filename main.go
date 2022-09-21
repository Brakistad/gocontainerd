// simple api setup
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong")
}

func main() {
	
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setup router
	router := mux.NewRouter()

	// setup routes
	router.HandleFunc("/api", ping).Methods("GET")

	// setup server
	srv := &http.Server{
		Handler: router,
		Addr:    "localhost" + ":" + os.Getenv("PORT"),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// start server
	log.Fatal(srv.ListenAndServe())
}
