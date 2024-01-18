package main

import (
	"fmt"
	"testing"
)

func TestAVeryBigSum(t *testing.T) {
	var tests = []struct {
		input []int64
		want  int64
	}{
		{[]int64{1, 2, 3, 4, 5, 6, 7, 8}, 36},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input: %v", test.input)

		t.Run(testName, func(t *testing.T) {
			res := aVeryBigSum(test.input)
			if res != test.want {
				t.Errorf("Expected %v but got %v", test.want, res)
			}
		})
	}
}
