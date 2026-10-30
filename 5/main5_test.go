package main

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	gotExists, gotValues := intersection(a, b)
	if !gotExists {
		t.Fatal("expected intersection")
	}
	wantValues := []int{64, 3}
	if !reflect.DeepEqual(gotValues, wantValues) {
		t.Fatalf("got %v, want %v", gotValues, wantValues)
	}
}

func TestIntersectionEmpty(t *testing.T) {
	gotExists, gotValues := intersection([]int{1, 2}, []int{3, 4})
	if gotExists || len(gotValues) != 0 {
		t.Fatalf("got (%v, %v), want (false, empty)", gotExists, gotValues)
	}
}
