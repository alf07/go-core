package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMergeChannels(t *testing.T) {
	first := make(chan int, 2)
	second := make(chan int, 2)
	first <- 1
	first <- 2
	second <- 3
	second <- 4
	close(first)
	close(second)

	got := make([]int, 0, 4)
	for value := range mergeChannels(first, second) {
		got = append(got, value)
	}

	want := []int{1, 2, 3, 4}
	sort.Ints(got)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func TestMergeNoChannels(t *testing.T) {
	out := mergeChannels()
	if _, ok := <-out; ok {
		t.Fatal("output should be closed")
	}
}
