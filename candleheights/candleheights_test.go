package main

import "testing"

func TestBirthdayCakeCandles(t *testing.T) {
	res := birthdayCakeCandles([]int32{3, 2, 1, 3})
	var expected int32 = 2

	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}
