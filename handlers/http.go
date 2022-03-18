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
	var request models.Human
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
	rw.Header().Set("Content-Type", "application/json")
	count_mutant_dna, count_human_dna, ratio := db.GetListDNA()
	stats := &models.Stats{
		CountMutantDna: count_mutant_dna,
		CountHumanDna: count_human_dna,
		Ratio: ratio,
	}
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(stats)
}
