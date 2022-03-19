package models

import (
	"fmt"
	matrixFunctions "isMutant/utils"
)

//Human Objeto persona
type Human struct {
	Dna []string `json:"dna"`
	Mutant bool `json:"isMutant"`
}

//IsMutant Método de Persona paa identificar si es mutante o no
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
	fmt.Println("CountTotalConsecutive: ", countTotalConsecutive)
	if countTotalConsecutive > 1 {
		h.Mutant = true
	} else {
		h.Mutant = false
	}
}

//Stats Objeto para las estadísticas
type Stats struct {
	CountMutantDna int     `json:"count_mutant_dna"`
	CountHumanDna  int     `json:"count_human_dna"`
	Ratio          float32 `json:"ratio"`
}