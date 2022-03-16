package handlers

import (
	"fmt"
	"net/http"
)

func CheckIsMutant(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Llega el ADN")
}

func GetStats(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Obtener estadísticas")
}
