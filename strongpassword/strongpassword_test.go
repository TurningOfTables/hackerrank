package main

import (
	"fmt"
	"testing"
)

func TestStrongPassword(t *testing.T) {

	var tests = []struct {
		currPassLen int32
		password    string
		want        int32
	}{
		{3, "Ab1", 3},
		{11, "#HackerRank", 1},
		{6, "Ab1hk%", 0},
		{1, "A", 5},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("%v,%v", test.currPassLen, test.password)
		t.Run(testName, func(t *testing.T) {
			res := minimumNumber(test.currPassLen, test.password)
			if res != test.want {
				t.Errorf("Expected %v but got %v", test.want, res)
			}
		})
	}
}
