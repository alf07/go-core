package main

import (
	"fmt"
	"sync"
)

func mergeChannels(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, channel := range channels {
		go func(input <-chan int) {
			defer wg.Done()
			for value := range input {
				out <- value
			}
		}(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	first := make(chan int)
	second := make(chan int)

	go func() {
		defer close(first)
		first <- 1
		first <- 2
	}()

	go func() {
		defer close(second)
		second <- 3
		second <- 4
	}()

	for value := range mergeChannels(first, second) {
		fmt.Println(value)
	}
}
