package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Routes Creación de rutas del proyecto
func Routes() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/v1/mutant/", CheckIsMutant).Methods("POST")
	mux.HandleFunc("/v1/stats/", GetStats).Methods("GET")
	return mux
}

//ServerListener Creacion de servidor
func ServerListener(port string)  {
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, Routes()))
}