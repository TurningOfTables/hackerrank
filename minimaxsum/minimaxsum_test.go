package main

import (
	"fmt"
	"testing"
)

func TestMiniMaxSum(t *testing.T) {
	var tests = []struct {
		input   []int32
		wantMin int64
		wantMax int64
	}{
		{[]int32{1, 3, 5, 7, 9}, 16, 24},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input %v, expected min %d, expected max %d", test.input, test.wantMin, test.wantMax)

		t.Run(testName, func(t *testing.T) {
			min, max := miniMaxSum(test.input)

			if min != test.wantMin || max != test.wantMax {
				t.Errorf("Expected min %d, max %d. Got min %d, max %d", test.wantMin, test.wantMax, min, max)
			}
		})
	}
}
