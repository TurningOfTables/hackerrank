package main

import (
	"fmt"
	"testing"
)

func TestReverseArray(t *testing.T) {
	var tests = []struct {
		input []int32
		want  []int32
	}{
		{input: []int32{1, 4, 3, 2}, want: []int32{2, 3, 4, 1}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.want), func(t *testing.T) {
			res := reverseArray(test.input)

			for k, v := range res {
				if v != test.want[k] {
					t.Errorf("Expected %v but got %v", test.want, res)
				}
			}
		})
	}
}
