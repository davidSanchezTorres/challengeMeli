package models

import (
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
		countTotalConsecutive += matrixFunctions.CountHorizontal(matrix[i])
		countTotalConsecutive += matrixFunctions.CountVertical(matrix, i)
		countTotalConsecutive += matrixFunctions.CountUpToDown1(matrix, i)
		countTotalConsecutive += matrixFunctions.CountUpToDown2(matrix, i)
		countTotalConsecutive += matrixFunctions.CountDownToUp1(matrix, i)
		countTotalConsecutive += matrixFunctions.CountDownToUp2(matrix, i)
	}

	if countTotalConsecutive > 1 {
		h.Mutant = true
	} else {
		h.Mutant = false
	}
}
