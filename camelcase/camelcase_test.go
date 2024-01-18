package main

import (
	"testing"
)

func TestCamelCase(t *testing.T) {
	res := camelcase("saveChangesInTheEditor")
	if res != 5 {
		t.Errorf("Expected 5 but got %v", res)
	}
}
