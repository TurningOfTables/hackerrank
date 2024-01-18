package main

import (
	"fmt"
	"testing"
)

func TestHandshake(t *testing.T) {

	var tests = []struct {
		input int32
		want  int32
	}{
		{1, 0},
		{2, 1},
		{3, 3},
		{4, 6},
		{5, 10},
		{6, 15},
		{10, 45},
		{100, 4950},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("input: %v", test.input), func(t *testing.T) {
			res := handshake(test.input)
			if res != test.want {
				t.Errorf("expected: %v got: %v", test.want, res)
			}
		})
	}
}
