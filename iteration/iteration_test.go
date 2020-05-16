package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	num := 4
	repeated := Repeated("a", num)
	expected := strings.Repeat("a", num)

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	num := 4
	for i := 0; i < b.N; i++ {
		Repeated("a", num)
	}
}

func ExampleRepeat() {
	num := 4
	expected := strings.Repeat("a", num)
	fmt.Println(expected)
	// Output: aaaa
}
