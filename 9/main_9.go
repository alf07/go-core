package main

import "fmt"

func numberPipeline(input <-chan uint8) <-chan float64 {
	output := make(chan float64)

	go func() {
		defer close(output)
		for value := range input {
			number := float64(value)
			output <- number * number * number
		}
	}()

	return output
}

func main() {
	input := make(chan uint8)
	output := numberPipeline(input)

	go func() {
		defer close(input)
		for _, value := range []uint8{1, 2, 3, 4} {
			input <- value
		}
	}()

	for value := range output {
		fmt.Println(value)
	}
}
