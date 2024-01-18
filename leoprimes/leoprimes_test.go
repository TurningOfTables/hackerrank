package main

import (
	"fmt"
	"testing"
)

func TestAllSolutions(t *testing.T) {
	var tests = map[string]func(int64) int64{
		"primeCount":  primeCount,
		"primeCount1": primeCount1,
		"primeCount2": primeCount2,
		"primeCount3": primeCount3,
	}

	var data = []struct {
		input int64
		want  int64
	}{
		{100, 3},
		{2, 1},
		{3, 1},
		{500, 4},
		{5000, 5},
		{10000000000, 10},
	}

	for name, function := range tests {
		for _, d := range data {
			testName := fmt.Sprintf("function: %v, input: %v", name, d)
			t.Run(testName, func(t *testing.T) {
				if function(d.input) != d.want {
					t.Errorf("expected %v, got %v", d.want, d.input)
				}
			})
		}

	}
}
