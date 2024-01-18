package main

import (
	"fmt"
	"testing"
)

func TestLowestTriangle(t *testing.T) {
	var tests = []struct {
		base int32
		area int32
		want int32
	}{
		{17, 100, 12},
		{8378, 42565, 11},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("base: %v, area: %v", test.base, test.area), func(t *testing.T) {
			res := lowestTriangle(test.base, test.area)

			if res != test.want {
				t.Errorf("expected: %v, got: %v", test.want, res)
			}
		})
	}
}
