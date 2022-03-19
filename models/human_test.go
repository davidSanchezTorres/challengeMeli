package models

import "testing"

func TestHumanIsMutant(t *testing.T) {
	type humanTest struct {
		Dna []string
	}
	dna1 := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"}
	human1 := humanTest{
		Dna: dna1,
	}
	dna2 := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG2"}
	human2 := humanTest{
		Dna: dna2,
	}
	dna3 := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG", "TCACTG"}
	human3 := humanTest{
		Dna: dna3,
	}
	tests := []struct {
		human  humanTest
		dna    []string
		mutant bool
	}{
		{human1, dna1, true},
		{human2, dna2, false},
		{human3, dna3, false},
	}
	for _, test := range tests {
		h := &Human{
			Dna: test.human.Dna,
		}
		h.IsMutant(test.dna)
		if h.Mutant != test.mutant {
			t.Errorf("Is mutant method incorrect = %v, expected: %v", h.Mutant, test.mutant)
		}
	}
}
