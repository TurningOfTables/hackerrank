package main

import "testing"

func TestStrangeCounter(t *testing.T) {
	res := strangeCounter(15)
	var expected int64 = 7

	if res != expected {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}
