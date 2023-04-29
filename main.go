package main

import (
	"log"
	"net/http"

	"mood-api/handlers"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/api/health", handlers.PingHandler)
	r.HandleFunc("/api/mood", handlers.MoodHandler)
	http.Handle("/", r)

	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
