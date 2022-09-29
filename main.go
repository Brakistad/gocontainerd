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
	fmt.Fprintf(w, "Pong !")
	fmt.Fprint(os.Stdout, "Pong !")
}

func main() {
	fmt.Fprint(os.Stdout, "Starting server...")
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file") 
		fmt.Fprint(os.Stdout, "Error loading .env file")
	}
	fmt.Fprint(os.Stdout, "Loaded env variables...")
	// setup router
	router := mux.NewRouter()

	// setup routes
	router.HandleFunc("/ping", ping).Methods("GET")

	fmt.Fprint(os.Stdout, "")

	// make channel for graceful shutdown
	c := make(chan os.Signal, 1)

	go func() {
		// setup server
		srv := &http.Server{
			Handler: router,
			Addr:    "0.0.0.0" + ":" + os.Getenv("PORT"),
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
	fmt.Fprint(os.Stdout, "Server started on port "+os.Getenv("PORT"))
	// wait for graceful shutdown
	go func() {
		<-c
		fmt.Fprint(os.Stdout, "Shutting down server...")
		os.Exit(0)
	}()
}