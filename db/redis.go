package db

import (
	"fmt"
)

func CreateDNA(dna []string, isMutant bool) {
	var dnaKey string
	for _, row := range dna {
		dnaKey += row + "+"
	}
	fmt.Printf("SET,'%v','%v'", dnaKey, isMutant)
}

func GetListDNS() {
	fmt.Println("Obtener lista de dns")
}
