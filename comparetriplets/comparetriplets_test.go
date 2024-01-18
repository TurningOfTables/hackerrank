package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCompareTriplets(t *testing.T) {
	var tests = []struct {
		a    []int32
		b    []int32
		want []int32
	}{
		{[]int32{5, 6, 7}, []int32{3, 6, 10}, []int32{1, 1}},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("Input: %v / %v", test.a, test.b)
		t.Run(testName, func(t *testing.T) {
			res := compareTriplets(test.a, test.b)
			if !reflect.DeepEqual(res, test.want) {
				t.Errorf("Expected %v but got %v", test.want, res)
			}
		})
	}
}
