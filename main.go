package main

import (
	"fmt"
	"isMutant/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	// Creación de endpoint
	mux.HandleFunc("/v1/mutant/", handlers.CheckIsMutant).Methods("POST")
	mux.HandleFunc("/v1/stats/", handlers.GetStats).Methods("GET")

	fmt.Println("Listener on :3000")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
