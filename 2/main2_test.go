package main

import (
	"reflect"
	"testing"
)

func TestSliceExample(t *testing.T) {
	got := sliceExample([]int{1, 2, 3, 4, 6, 7})
	want := []int{2, 4, 6}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestAddElementsDoesNotModifyOriginal(t *testing.T) {
	original := []int{1, 2, 3}
	got := addElements(original, 4)

	if !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Fatalf("got %v", got)
	}
	if !reflect.DeepEqual(original, []int{1, 2, 3}) {
		t.Fatalf("original was modified: %v", original)
	}
}

func TestCopySliceIsIndependent(t *testing.T) {
	original := []int{1, 2, 3}
	copyOfOriginal := copySlice(original)
	original[0] = 99

	if copyOfOriginal[0] != 1 {
		t.Fatalf("copy changed with original: %v", copyOfOriginal)
	}
}

func TestRemoveElement(t *testing.T) {
	got := removeElement([]int{1, 2, 3, 4}, 2)
	want := []int{1, 2, 4}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestRemoveElementInvalidIndex(t *testing.T) {
	original := []int{1, 2, 3}
	got := removeElement(original, 10)
	if !reflect.DeepEqual(got, original) {
		t.Fatalf("got %v, want %v", got, original)
	}
}
