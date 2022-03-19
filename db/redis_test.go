package db

import (
	"testing"
)

func TestGetListDNA(t *testing.T) {
	tests := []struct {
		name        string
		mutantCount int
		humanCount  int
		ratio       float32
	}{
		{"Test 1", 1, 1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			muntantC, humanC, ratio := GetListDNA()
			if muntantC >= tt.mutantCount {
				t.Errorf("GetListDNA() muntantC = %v, mutantCount %v", muntantC, tt.mutantCount)
			}
			if humanC >= tt.humanCount {
				t.Errorf("GetListDNA() humanC = %v, mutantCount %v", humanC, tt.humanCount)
			}
			if ratio >= tt.ratio {
				t.Errorf("GetListDNA() ratio = %v, mutantCount %v", ratio, tt.ratio)
			}
		})
	}
}
