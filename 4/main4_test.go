package main

import (
	"reflect"
	"testing"
)

func TestDifference(t *testing.T) {
	first := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	second := []string{"banana", "date", "fig"}
	want := []string{"apple", "cherry", "43", "lead", "gno1"}

	if got := difference(first, second); !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestDifferenceEmpty(t *testing.T) {
	got := difference([]string{"a", "b"}, []string{"a", "b"})
	if len(got) != 0 {
		t.Fatalf("got %v, want empty slice", got)
	}
}
