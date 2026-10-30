package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func randomGenerator(ctx context.Context, min, max int) <-chan int {
	out := make(chan int)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	go func() {
		defer close(out)
		for {
			value := r.Intn(max-min+1) + min
			select {
			case out <- value:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	for value := range randomGenerator(ctx, 1, 100) {
		fmt.Println(value)
	}
}
