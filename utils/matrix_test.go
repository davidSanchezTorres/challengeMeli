package utils

import (
	"testing"
)

func TestConvertListToMatrix(t *testing.T) {
	dna1 := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	res1 := [][]string{
		{"A", "T", "G", "C", "G", "A"},
		{"C", "A", "G", "T", "G", "C"},
		{"T", "T", "A", "T", "G", "T"},
		{"A", "G", "A", "A", "G", "G"},
		{"C", "C", "C", "C", "T", "A"},
		{"T", "C", "A", "C", "T", "G"},
	}
	dna2 := []string{"ATGCGAT", "CAGTGCG", "TTATGTA", "AGAAGGA", "CCCCTAC", "TCACTGC", "CTCAGTC"}
	res2 := [][]string{
		{"A", "T", "G", "C", "G", "A", "T"},
		{"C", "A", "G", "T", "G", "C", "G"},
		{"T", "T", "A", "T", "G", "T", "A"},
		{"A", "G", "A", "A", "G", "G", "A"},
		{"C", "C", "C", "C", "T", "A", "C"},
		{"T", "C", "A", "C", "T", "G", "C"},
		{"C", "T", "C", "A", "G", "T", "C"},
	}
	tests := []struct {
		row    []string
		matrix [][]string
		size   int
		err    bool
	}{
		{dna1, res1, 6, false},
		{dna2, res2, 7, false},
	}

	for _, test := range tests {
		matrix, size, err := ConvertListToMatrix(test.row)
		if (err != nil) != test.err {
			t.Errorf("ConvertListToMatrix() error = %v, err %v", err, test.err)
			return
		}
		if size != test.size {
			t.Errorf("Size response incorrect = %v, err %v", size, test.size)
		}
		if len(matrix) != len(test.matrix) {
			t.Errorf("Size matrix incorrect = %v, err %v", len(matrix), len(test.matrix))
		}
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if matrix[i][j] != test.matrix[i][j] {
					t.Errorf("Value incorrect = %v = %v (%v,%v)", matrix[i][j], test.matrix[i][j], i, j)
				}
			}
		}
	}
}

// func TestConvertListToMatrix(t *testing.T) {
// 	type args struct {
// 		row []string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    [][]string
// 		want1   int
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1, err := ConvertListToMatrix(tt.args.row)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("ConvertListToMatrix() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("ConvertListToMatrix() got = %v, want %v", got, tt.want)
// 			}
// 			if got1 != tt.want1 {
// 				t.Errorf("ConvertListToMatrix() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }
