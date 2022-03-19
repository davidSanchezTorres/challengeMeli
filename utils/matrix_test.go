package utils

import (
	"testing"
)

func TestCounTSecuenSuccess(t *testing.T) {
	dna1 := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	res1 := [][]string{
		{"A", "T", "G", "C", "G", "A"},
		{"C", "A", "G", "T", "G", "C"},
		{"T", "T", "A", "T", "G", "T"},
		{"A", "G", "A", "A", "G", "G"},
		{"C", "C", "C", "C", "T", "A"},
		{"T", "C", "A", "C", "T", "G"},
	}

	dna11 := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TTTTCG"}
	res11 := [][]string{
		{"A", "T", "G", "C", "G", "A"},
		{"C", "A", "G", "T", "G", "C"},
		{"T", "T", "A", "T", "G", "T"},
		{"A", "G", "A", "A", "G", "G"},
		{"C", "C", "C", "C", "T", "A"},
		{"T", "T", "T", "T", "C", "G"},
	}
	dna2 := []string{"ATGCGAT", "CAGTGCG", "TTATGTA", "AGAAGGC", "CCCCTAC", "TCACTGC", "CTCAGTC"}
	res2 := [][]string{
		{"A", "T", "G", "C", "G", "A", "T"},
		{"C", "A", "G", "T", "G", "C", "G"},
		{"T", "T", "A", "T", "G", "T", "A"},
		{"A", "G", "A", "A", "G", "G", "C"},
		{"C", "C", "C", "C", "T", "A", "C"},
		{"T", "C", "A", "C", "T", "G", "C"},
		{"C", "T", "C", "A", "G", "T", "C"},
	}
	dna3 := []string{"ATGCGATGT", "TTTGGCGCT", "TTGGGGCCT", "TCCTGTAAC", "CCCCTTTAC", "TTTTCCGAA", "TTGGCCCGA", "CGGGTCCCC", "AAAATTCGG"}
	res3 := [][]string{
		{"A", "T", "G", "C", "G", "A", "T", "G", "T"},
		{"T", "T", "T", "G", "G", "C", "G", "C", "T"},
		{"T", "T", "G", "G", "G", "G", "C", "C", "T"},
		{"T", "C", "C", "T", "G", "T", "A", "A", "C"},
		{"C", "C", "C", "C", "T", "T", "T", "A", "C"},
		{"T", "T", "T", "T", "C", "C", "G", "A", "A"},
		{"T", "T", "G", "G", "C", "C", "C", "G", "A"},
		{"C", "G", "G", "G", "T", "C", "C", "C", "C"},
		{"A", "A", "A", "A", "T", "T", "C", "G", "G"},
	}
	dna31 := []string{"ATGCGATGT", "TTTGGCGCT", "TTGGGGCCT", "TTCTGTAAC", "CCCCTTTAC", "TTTTCCGAA", "TTGGCCCGA", "CGGGTCCCC", "AAAATTCGG"}
	res31 := [][]string{
		{"A", "T", "G", "C", "G", "A", "T", "G", "T"},
		{"T", "T", "T", "G", "G", "C", "G", "C", "T"},
		{"T", "T", "G", "G", "G", "G", "C", "C", "T"},
		{"T", "T", "C", "T", "G", "T", "A", "A", "C"},
		{"C", "C", "C", "C", "T", "T", "T", "A", "C"},
		{"T", "T", "T", "T", "C", "C", "G", "A", "A"},
		{"T", "T", "G", "G", "C", "C", "C", "G", "A"},
		{"C", "G", "G", "G", "T", "C", "C", "C", "C"},
		{"A", "A", "A", "A", "T", "T", "C", "G", "G"},
	}
	dna32 := []string{"ATGCGATGT", "TTTGGCGCT", "TTGGGGCCT", "TTCTGTAAC", "CCCCTTTAC", "TTTTCCGAA", "TTGGCCCGA", "CCCCTCCCC", "AAAATTCGG"}
	res32 := [][]string{
		{"A", "T", "G", "C", "G", "A", "T", "G", "T"},
		{"T", "T", "T", "G", "G", "C", "G", "C", "T"},
		{"T", "T", "G", "G", "G", "G", "C", "C", "T"},
		{"T", "T", "C", "T", "G", "T", "A", "A", "C"},
		{"C", "C", "C", "C", "T", "T", "T", "A", "C"},
		{"T", "T", "T", "T", "C", "C", "G", "A", "A"},
		{"T", "T", "G", "G", "C", "C", "C", "G", "A"},
		{"C", "C", "C", "C", "T", "C", "C", "C", "C"},
		{"A", "A", "A", "A", "T", "T", "C", "G", "G"},
	}
	tests := []struct {
		row               []string
		matrix            [][]string
		size              int
		err               bool
		countHorizontal   int
		countVertical     int
		countUpToDown1    int
		countUpToDown2    int
		countDownToUp1    int
		countDownToUp2    int
		countSecuenMatrix int
	}{
		{dna1, res1, 6, false, 1, 1, 1, 0, 0, 0, 3},
		{dna11, res11, 6, false, 2, 1, 1, 0, 0, 0, 4},
		{dna2, res2, 7, false, 1, 2, 1, 0, 0, 0, 4},
		{dna3, res3, 9, false, 5, 1, 1, 0, 1, 0, 8},
		{dna31, res31, 9, false, 5, 2, 1, 0, 1, 0, 9},
		{dna32, res32, 9, false, 6, 2, 1, 0, 1, 0, 10},
	}

	for _, test := range tests {
		matrix, size, _ := ConvertListToMatrix(test.row)
		if size != test.size {
			t.Errorf("Size response incorrect = %v, expected: %v", size, test.size)
		}
		if len(matrix) != len(test.matrix) {
			t.Errorf("Size matrix incorrect = %v, expected: %v", len(matrix), len(test.matrix))
		}
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if matrix[i][j] != test.matrix[i][j] {
					t.Errorf("Value incorrect = %v expected: %v (%v,%v)", matrix[i][j], test.matrix[i][j], i, j)
				}
			}
		}

		var countHorizontal int
		for i := 0; i < len(matrix); i++ {
			countHorizontal += CountHorizontal(matrix[i])
		}
		if countHorizontal != test.countHorizontal {
			t.Errorf("Count Horizontal incorrect = %v, expected: %v", countHorizontal, test.countHorizontal)
		}

		var countVertical int
		for i := 0; i < len(matrix); i++ {
			countVertical += CountVertical(matrix, i)
		}
		if countVertical != test.countVertical {
			t.Errorf("Count Vertical incorrect = %v, expected: %v", countVertical, test.countVertical)
		}

		var countUpToDown1 int
		for i := 0; i < len(matrix); i++ {
			countUpToDown1 += CountUpToDown1(matrix, i)
		}
		if countUpToDown1 != test.countUpToDown1 {
			t.Errorf("Count countUpToDown1 incorrect = %v, expected: %v", countUpToDown1, test.countUpToDown1)
		}

		var countUpToDown2 int
		for i := 0; i < len(matrix); i++ {
			countUpToDown2 += CountUpToDown2(matrix, i)
		}
		if countUpToDown2 != test.countUpToDown2 {
			t.Errorf("Count countUpToDown2 incorrect = %v, expected: %v", countUpToDown2, test.countUpToDown2)
		}

		var countDownToUp1 int
		for i := 0; i < len(matrix); i++ {
			countDownToUp1 += CountDownToUp1(matrix, i)
		}
		if countDownToUp1 != test.countDownToUp1 {
			t.Errorf("Count countDownToUp1 incorrect = %v, expected: %v", countDownToUp1, test.countDownToUp1)
		}

		var countDownToUp2 int
		for i := 0; i < len(matrix); i++ {
			countDownToUp2 += CountDownToUp2(matrix, i)
		}
		if countDownToUp2 != test.countDownToUp2 {
			t.Errorf("Count countDownToUp2 incorrect = %v, expected: %v", countDownToUp2, test.countDownToUp2)
		}
	}
}

func TestCounTSecuenErrorRegex(t *testing.T) {
	dna1Regex := []string{"A", "T", "G", "C", "G", "A"}
	dna2Regex := []string{"A", "H", "G", "C", "G", "A"}
	dna3Regex := []string{"A", "G", "*", "C", "G", "A"}
	dna4Regex := []string{"A", "G", "T", "5", "G", "A"}
	dna5Regex := []string{"A", "", "T", "C", "G", "A"}

	tests := []struct {
		row        []string
		validRegex bool
	}{
		{dna1Regex, true},
		{dna2Regex, false},
		{dna3Regex, false},
		{dna4Regex, false},
		{dna5Regex, false},
	}

	for _, test := range tests {
		valid := validateRow(test.row)
		if valid != test.validRegex {
			t.Errorf("ValidateRow incorrect = %v, expected: %v ", test.row, valid)
		}
	}
}

func TestCounTSecuenError(t *testing.T) {
	dna1 := []string{"A*GCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	dna2 := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG2"}
	dna3 := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG", "TCACTA"}
	dna4 := []string{"AHTGCG", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}

	tests := []struct {
		row []string
		err bool
	}{
		{dna1, true},
		{dna2, true},
		{dna3, true},
		{dna4, true},
	}

	for _, test := range tests {
		_, _, err := ConvertListToMatrix(test.row)
		if (err != nil) && !test.err {
			t.Errorf("Error converting to array = %v, expected: %v", (err != nil), test.err)
		}
	}
}

// go test
// go test -coverprofile="coverage.out"
// go tool cover --func=coverage.out
// go tool cover --html=coverage.out
