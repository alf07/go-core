package main

import (
	"testing"
)

func TestTypeNames(t *testing.T) {
	got := typeNames(42, 3.14, "Golang", true, complex64(1+2i))
	want := []string{"int", "float64", "string", "bool", "complex64"}

	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("index %d: got %q, want %q", i, got[i], want[i])
		}
	}
}

func TestConcatenateValues(t *testing.T) {
	got := concatenateValues(42, 052, 0x2A, 3.14, "Golang", true, complex64(1+2i))
	want := "4242423.14Golangtrue(1+2i)"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

func TestHashRunesWithSalt(t *testing.T) {
	got := hashRunesWithSalt([]rune("abcdef"), salt)
	want := hashRunesWithSalt([]rune("abc"+salt+"def"), "")

	if got != want {
		t.Fatalf("unexpected hash: got %x, want %x", got, want)
	}
}
