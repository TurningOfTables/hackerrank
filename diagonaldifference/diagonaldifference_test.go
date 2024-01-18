package main

import (
	"fmt"
	"testing"
)

func TestDiagonalDif(t *testing.T) {
	var tests = []struct {
		matrix [][]int32
		want   int32
	}{
		{matrix: [][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}, want: 15},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Matrix: %v, Expected: %d", test.matrix, test.want)

		t.Run(testName, func(t *testing.T) {
			res := diagonalDifference(test.matrix)

			if res != test.want {
				t.Errorf("Expected %v, Got %v", test.want, res)
			}
		})
	}
}
