package handlers

import (
	"encoding/json"
	"fmt"
	"isMutant/models"
	"net/http"
)

func CheckIsMutant(rw http.ResponseWriter, r *http.Request) {
	var request models.DnaRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		fmt.Fprintln(rw, "Obtener estadísticas")
	}
	human := &models.Human{}
	human.IsMutant(request.Dna)
	fmt.Fprintln(rw, human.Mutant)
}

func GetStats(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Obtener estadísticas")
}
