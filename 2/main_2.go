package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sliceExample(numbers []int) []int {
	result := make([]int, 0, len(numbers))
	for _, number := range numbers {
		if number%2 == 0 {
			result = append(result, number)
		}
	}
	return result
}

func addElements(numbers []int, number int) []int {
	result := make([]int, len(numbers), len(numbers)+1)
	copy(result, numbers)
	result = append(result, number)
	return result
}

func copySlice(numbers []int) []int {
	result := make([]int, len(numbers))
	copy(result, numbers)
	return result
}

func removeElement(numbers []int, index int) []int {
	if index < 0 || index >= len(numbers) {
		return copySlice(numbers)
	}

	result := make([]int, 0, len(numbers)-1)
	result = append(result, numbers[:index]...)
	result = append(result, numbers[index+1:]...)
	return result
}

func randomSlice(size int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make([]int, size)
	for i := range result {
		result[i] = r.Intn(100)
	}
	return result
}

func main() {
	originalSlice := randomSlice(10)
	fmt.Println("Original:", originalSlice)
	fmt.Println("Even:", sliceExample(originalSlice))
	fmt.Println("Added:", addElements(originalSlice, 100))
	fmt.Println("Copy:", copySlice(originalSlice))
	fmt.Println("Without index 2:", removeElement(originalSlice, 2))
}
