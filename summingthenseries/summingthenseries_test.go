package main

import (
	"testing"
)

func TestSummingSeries(t *testing.T) {
	res := summingSeries(5351871996120528)
	expected := int32(578351320)

	if res != expected {
		t.Errorf("Expected %v but got %v", expected, res)
	}
}

func BenchmarkSummingSeries(b *testing.B) {
	for x := 0; x <= b.N; x++ {
		summingSeries(5351871996120528)
	}
}
