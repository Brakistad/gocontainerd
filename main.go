// simple api setup
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"sync"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var wg sync.WaitGroup

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong !")
	fmt.Println("Pong !")
}

func shutdown(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Shutting down server...")
	wg.Done()
}

func main() {
	fmt.Println("Starting server...")
	wg.Add(1)

	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file") 
		fmt.Println("Error loading .env file")
	}
	fmt.Println("Loaded env variables...")
	// setup router
	router := mux.NewRouter()

	// setup routes
	router.HandleFunc("/ping", ping).Methods("GET")
	router.HandleFunc("/shutdown", shutdown).Methods("GET")

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
	fmt.Println("Server started on port "+os.Getenv("PORT"))
	wg.Wait()
	fmt.Println("Server stopped")
}