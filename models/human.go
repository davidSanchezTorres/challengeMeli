package models

import (
	"fmt"
	"strings"
)

type Human struct {
	Mutant bool `json:"isMutant"`
}

func convertListToMatrix(row []string) ([][]string, int, error) {
	matrix := [][]string{}
	for _, base := range row {
		row := strings.Split(base, "")
		matrix = append(matrix, row)
	}
	return matrix, len(matrix), nil
}

func findConsecutive(row []string) int {
	var consecutives int
	countTemp := 1
	for i := 0; i < len(row); i++ {
		if i < len(row)-1 { // i < len(row) para evitar comparar la última con la siguiente que no existe
			if row[i] == row[i+1] {
				countTemp++
				if countTemp == 4 {
					consecutives++
					countTemp = 0
					// fmt.Println("Encontró en: ", row)
				}
			} else {
				countTemp = 0
			}
		}
	}
	return consecutives
}

func (h *Human) IsMutant(dna []string) {
	matrix, sizeMatrix, err := convertListToMatrix(dna)
	if err != nil {
		h.Mutant = false
	}
	countTotalConsecutive := 0
	for i := 0; i < sizeMatrix; i++ {
		countTotalConsecutive += findConsecutive(matrix[i])
		//Recorrido para obtener Verticales
		var rowCount []string
		for j := 0; j < sizeMatrix; j++ {
			rowCount = append(rowCount, matrix[j][i])
		}
		countTotalConsecutive += findConsecutive(rowCount)

		//Recorrido para obtener diagonales de arriba hacia abajo desde la mitad
		jTemp := 0
		rowCount = nil
		for iDiag := i; iDiag < sizeMatrix; iDiag++ {
			rowCount = append(rowCount, matrix[iDiag][jTemp])
			jTemp++
		}
		countTotalConsecutive += findConsecutive(rowCount)

		//Recorrido para obtener diagonales de arriba hacia abajo desde uno despupes de la mitad hacia la derecha
		jTemp = 0
		rowCount = nil
		for jDiag := i + 1; jDiag < sizeMatrix; jDiag++ {
			rowCount = append(rowCount, matrix[jTemp][jDiag])
			jTemp++
		}
		countTotalConsecutive += findConsecutive(rowCount)

		//Recorrido para obtener diagonales de abajo hacia arriba desde la mitad
		jTemp = 0
		rowCount = nil
		for iDiag := (sizeMatrix - 1) - i; iDiag >= 0; iDiag-- {
			rowCount = append(rowCount, matrix[iDiag][jTemp])
			jTemp++
		}
		countTotalConsecutive += findConsecutive(rowCount)

		//Recorrido para obtener diagonales de abajo hacia arriba desde uno después de la mitad
		jTemp = i + 1
		rowCount = nil
		for iDiag := sizeMatrix - 1; iDiag >= (i + 1); iDiag-- {
			rowCount = append(rowCount, matrix[iDiag][jTemp])
			jTemp++
		}
		countTotalConsecutive += findConsecutive(rowCount)
	}

	fmt.Println("countTotalConsecutive:", countTotalConsecutive)
	if countTotalConsecutive > 1 {
		h.Mutant = true
	} else {
		h.Mutant = false
	}
}
