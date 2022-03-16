package main

import (
	"errors"
	"fmt"
	"strings"
)

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
		if i < len(row)-1 {
			if row[i] == row[i+1] {
				countTemp++
				// fmt.Printf("%s == %s: poicion(%d) asiertos: %d \n", row[i], row[i+1], i, consecutives)
				if countTemp == 4 {
					consecutives++
					countTemp = 0
					fmt.Println("Encontró en: ", row)
				}
			} else {
				countTemp = 0
			}
		}
	}
	return consecutives
}

func isMutant(dna []string) (bool, error) {
	matrix, sizeMatrix, err := convertListToMatrix(dna)
	if err != nil {
		return false, errors.New("Error to converter on matrix")
	}
	countTotalConsecutive := 0
	for i := 0; i < sizeMatrix; i++ {
		//Horizontales
		countTotalConsecutive += findConsecutive(matrix[i])
		//Verticales
		var column []string
		for j := 0; j < sizeMatrix; j++ {
			column = append(column, matrix[j][i])
		}
		countTotalConsecutive += findConsecutive(column)
	}

	// Diagonales hacia abajo
	for i := 0; i < sizeMatrix; i++ {
		jTemp := 0
		var diagonalDown []string
		for iDiag := i; iDiag < sizeMatrix; iDiag++ {
			diagonalDown = append(diagonalDown, matrix[iDiag][jTemp])
			jTemp++
		}
		jTemp = 0
		var diagonalUp []string
		for jDiag := i + 1; jDiag < sizeMatrix; jDiag++ {
			diagonalUp = append(diagonalUp, matrix[jTemp][jDiag])
			jTemp++
		}
		countTotalConsecutive += findConsecutive(diagonalDown)
		countTotalConsecutive += findConsecutive(diagonalUp)
	}

	fmt.Println("countTotalConsecutive:", countTotalConsecutive)
	if countTotalConsecutive > 1 {
		return true, nil
	} else {
		return false, nil
	}
}

func main() {
	dna := []string{
		"ATGCGASP",
		"CAGGGCSP",
		"TTATGTOP",
		"AGAAGGOP",
		"CCCCTAOP",
		"TCACTGOP",
		"TCCCTGSP",
		"FLACYGSP",
	}
	fmt.Println("1111111111")
	fmt.Println(dna)
	fmt.Println("2222222222")
	isMutant, err := isMutant(dna)
	if err != nil {
		fmt.Println("Ocurrió un error")
	}
	fmt.Println("Mutante: ", isMutant)
}
