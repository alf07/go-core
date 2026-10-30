package main

import (
	"fmt"
)

type WaitGroup struct {
	semaphore chan struct{}
	done      chan struct{}
	count     int
}

func NewWaitGroup() *WaitGroup {
	semaphore := make(chan struct{}, 1)
	semaphore <- struct{}{}
	return &WaitGroup{
		semaphore: semaphore,
		done:      make(chan struct{}),
	}
}

func (w *WaitGroup) Add(delta int) {
	<-w.semaphore
	w.count += delta
	if w.count < 0 {
		w.count = 0
		w.semaphore <- struct{}{}
		panic("negative WaitGroup counter")
	}
	if w.count == 0 {
		select {
		case <-w.done:
		default:
			close(w.done)
		}
	}
	w.semaphore <- struct{}{}
}

func (w *WaitGroup) Done() {
	w.Add(-1)
}

func (w *WaitGroup) Wait() {
	<-w.done
}

func main() {
	wg := NewWaitGroup()
	wg.Add(2)

	go func() {
		defer wg.Done()
		fmt.Println("worker 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("worker 2")
	}()

	wg.Wait()
	fmt.Println("all workers finished")
}
