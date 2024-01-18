package main

import "testing"

func TestTimeConversion(t *testing.T) {
	res := timeConversion("07:05:45PM")
	expected := "19:05:45"

	if res != expected {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}
