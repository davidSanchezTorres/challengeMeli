package utils

import (
	"fmt"
	"strings"
)

func ConvertListToMatrix(row []string) ([][]string, int, error) {
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
					fmt.Println("Encontró secuencia: ", row)
				}
			} else {
				countTemp = 0
			}
		}
	}
	return consecutives
}

func CountHorizontal(row []string) int {
	return findConsecutive(row)
}

func CountVertical(matrix [][]string, i int) int {
	var rowCount []string
	for j := 0; j < len(matrix); j++ {
		rowCount = append(rowCount, matrix[j][i])
	}
	return findConsecutive(rowCount)
}

//Recorrido para obtener diagonales de arriba hacia abajo desde la mitad
func CountUpToDown1(matrix [][]string, i int) int {
	jTemp := 0
	var rowCount []string
	for iDiag := i; iDiag < len(matrix); iDiag++ {
		rowCount = append(rowCount, matrix[iDiag][jTemp])
		jTemp++
	}
	return findConsecutive(rowCount)
}

//Recorrido para obtener diagonales de arriba hacia abajo desde uno despupes de la mitad hacia la derecha
func CountUpToDown2(matrix [][]string, i int) int {
	jTemp := 0
	var rowCount []string
	for jDiag := i + 1; jDiag < len(matrix); jDiag++ {
		rowCount = append(rowCount, matrix[jTemp][jDiag])
		jTemp++
	}
	return findConsecutive(rowCount)
}

//Recorrido para obtener diagonales de abajo hacia arriba desde la mitad
func CountDownToUp1(matrix [][]string, i int) int {
	jTemp := 0
	var rowCount []string
	for iDiag := (len(matrix) - 1) - i; iDiag >= 0; iDiag-- {
		rowCount = append(rowCount, matrix[iDiag][jTemp])
		jTemp++
	}
	return findConsecutive(rowCount)
}

//Recorrido para obtener diagonales de abajo hacia arriba desde uno después de la mitad
func CountDownToUp2(matrix [][]string, i int) int {
	jTemp := i + 1
	var rowCount []string
	for iDiag := len(matrix) - 1; iDiag >= (i + 1); iDiag-- {
		rowCount = append(rowCount, matrix[iDiag][jTemp])
		jTemp++
	}
	return findConsecutive(rowCount)
}
