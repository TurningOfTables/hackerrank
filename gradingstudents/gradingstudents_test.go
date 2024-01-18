package main

import "testing"

func TestGradingStudents(t *testing.T) {
	res := gradingStudents([]int32{73, 67, 38, 33})
	expected := []int32{75, 67, 40, 33}

	if len(res) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, res)
	} else {
		for k := range expected {
			if res[k] != expected[k] {
				t.Errorf("Expected %v, got %v", expected, res)
			}
		}
	}
}
