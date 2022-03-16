package utils

import "strings"

func ConvertListToMatrix(row []string) ([][]string, int, error) {
	matrix := [][]string{}
	for _, base := range row {
		row := strings.Split(base, "")
		matrix = append(matrix, row)
	}
	return matrix, len(matrix), nil
}

func FindConsecutive(row []string) int {
	var consecutives int
	countTemp := 1
	for i := 0; i < len(row); i++ {
		if i < len(row)-1 { // i < len(row) para evitar comparar la última con la siguiente que no existe
			if row[i] == row[i+1] {
				countTemp++
				if countTemp == 4 {
					consecutives++
					countTemp = 0
				}
			} else {
				countTemp = 0
			}
		}
	}
	return consecutives
}
