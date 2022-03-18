package handlers

import (
	"encoding/json"
	"fmt"
	"isMutant/db"
	"isMutant/models"
	"net/http"
)

func CheckIsMutant(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var request models.DnaRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(rw, "Error decoder body")
	}
	human := &models.Human{}
	human.IsMutant(request.Dna)
	if human.Mutant {
		db.CreateDNA(request.Dna, true)
		rw.WriteHeader(http.StatusOK)
	} else {
		db.CreateDNA(request.Dna, false)
		rw.WriteHeader(http.StatusForbidden)
	}
}

func GetStats(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Obtener estadísticas")
}
