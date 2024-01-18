package main

import (
	"fmt"
	"testing"
)

func TestLeftRotation(t *testing.T) {

	var tests = []struct {
		rotations int32
		arr       []int32
		want      []int32
	}{
		{
			rotations: 7,
			arr:       []int32{1, 2, 3, 4, 5},
			want:      []int32{3, 4, 5, 1, 2},
		},
		{
			rotations: 6,
			arr:       []int32{1, 2, 3, 4, 5},
			want:      []int32{2, 3, 4, 5, 1},
		},
		{
			rotations: 137,

			arr:  []int32{1, 2, 3, 4, 5},
			want: []int32{3, 4, 5, 1, 2},
		},
		{
			rotations: 109,
			arr:       []int32{431, 397, 149, 275, 556, 362, 852, 789, 601, 357, 516, 575, 670, 507, 127, 888, 284, 405, 806, 27, 495, 879, 976, 467, 342, 356, 908, 750, 769, 947, 425, 643, 754, 396, 653, 595, 108, 75, 347, 394, 935, 252, 683, 966, 553, 724, 629, 567, 93, 494, 693, 965, 328, 187, 728, 389, 70, 288, 509, 252, 449},
			want:      []int32{93, 494, 693, 965, 328, 187, 728, 389, 70, 288, 509, 252, 449, 431, 397, 149, 275, 556, 362, 852, 789, 601, 357, 516, 575, 670, 507, 127, 888, 284, 405, 806, 27, 495, 879, 976, 467, 342, 356, 908, 750, 769, 947, 425, 643, 754, 396, 653, 595, 108, 75, 347, 394, 935, 252, 683, 966, 553, 724, 629, 567}},
	}

	var functions = map[string]func(int32, []int32) []int32{"rotateLeft": rotateLeft, "rotateLeftSlow": rotateLeftSlow}

	for fname, f := range functions {
		for _, test := range tests {
			testName := fmt.Sprintf("Function: %v, input length %v, rotations %v", fname, len(test.arr), test.rotations)
			t.Run(testName, func(t *testing.T) {
				res := f(test.rotations, test.arr)

				if len(res) != len(test.want) {
					t.Errorf("Expected %v but got %v", test.want, res)
				} else {
					for k := range test.want {
						if res[k] != test.want[k] {
							t.Errorf("Expected %v but got %v", test.want, res)
						}
					}

				}
			})
		}
	}

}

func BenchmarkLeftRotation(b *testing.B) {
	var tests = []struct {
		Name     string
		Function func(int32, []int32) []int32
	}{
		{
			Name:     "rotateLeft",
			Function: rotateLeft,
		},
		{
			Name:     "rotateLeftSlow",
			Function: rotateLeftSlow,
		},
	}

	var rotations int32 = 1200
	var arr = []int32{431, 397, 149, 275, 556, 362, 852, 789, 601, 357, 516, 575, 670, 507, 127, 888, 284, 405, 806, 27, 495, 879, 976, 467, 342, 356, 908, 750, 769, 947, 425, 643, 754, 396, 653, 595, 108, 75, 347, 394, 935, 252, 683, 966, 553, 724, 629, 567, 93, 494, 693, 965, 328, 187, 728, 389, 70, 288, 509, 252, 449}

	for _, test := range tests {
		b.Run(test.Name, func(b *testing.B) {
			for x := 0; x <= b.N; x++ {
				test.Function(rotations, arr)
			}
		})
	}
}
