package main

import (
	"reflect"
	"testing"
)

func TestNumberPipeline(t *testing.T) {
	input := make(chan uint8)
	output := numberPipeline(input)

	go func() {
		defer close(input)
		for _, value := range []uint8{1, 2, 3, 4} {
			input <- value
		}
	}()

	got := make([]float64, 0, 4)
	for value := range output {
		got = append(got, value)
	}

	want := []float64{1, 8, 27, 64}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
