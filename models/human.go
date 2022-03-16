package models

import (
	"fmt"
	matrixFunctions "isMutant/utils"
)

type Human struct {
	Mutant bool `json:"isMutant"`
}

func (h *Human) IsMutant(dna []string) {
	matrix, sizeMatrix, err := matrixFunctions.ConvertListToMatrix(dna)
	if err != nil {
		h.Mutant = false
	}
	countTotalConsecutive := 0
	for i := 0; i < sizeMatrix; i++ {
		countTotalConsecutive += matrixFunctions.FindConsecutive(matrix[i])
		//Recorrido para obtener Verticales
		var rowCount []string
		for j := 0; j < sizeMatrix; j++ {
			rowCount = append(rowCount, matrix[j][i])
		}
		countTotalConsecutive += matrixFunctions.FindConsecutive(rowCount)

		//Recorrido para obtener diagonales de arriba hacia abajo desde la mitad
		jTemp := 0
		rowCount = nil
		for iDiag := i; iDiag < sizeMatrix; iDiag++ {
			rowCount = append(rowCount, matrix[iDiag][jTemp])
			jTemp++
		}
		countTotalConsecutive += matrixFunctions.FindConsecutive(rowCount)

		//Recorrido para obtener diagonales de arriba hacia abajo desde uno despupes de la mitad hacia la derecha
		jTemp = 0
		rowCount = nil
		for jDiag := i + 1; jDiag < sizeMatrix; jDiag++ {
			rowCount = append(rowCount, matrix[jTemp][jDiag])
			jTemp++
		}
		countTotalConsecutive += matrixFunctions.FindConsecutive(rowCount)

		//Recorrido para obtener diagonales de abajo hacia arriba desde la mitad
		jTemp = 0
		rowCount = nil
		for iDiag := (sizeMatrix - 1) - i; iDiag >= 0; iDiag-- {
			rowCount = append(rowCount, matrix[iDiag][jTemp])
			jTemp++
		}
		countTotalConsecutive += matrixFunctions.FindConsecutive(rowCount)

		//Recorrido para obtener diagonales de abajo hacia arriba desde uno después de la mitad
		jTemp = i + 1
		rowCount = nil
		for iDiag := sizeMatrix - 1; iDiag >= (i + 1); iDiag-- {
			rowCount = append(rowCount, matrix[iDiag][jTemp])
			jTemp++
		}
		countTotalConsecutive += matrixFunctions.FindConsecutive(rowCount)
	}

	fmt.Println("countTotalConsecutive:", countTotalConsecutive)
	if countTotalConsecutive > 1 {
		h.Mutant = true
	} else {
		h.Mutant = false
	}
}
