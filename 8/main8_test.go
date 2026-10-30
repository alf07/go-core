package main

import (
	"testing"
	"time"
)

func TestWaitGroupWaitsForWorkers(t *testing.T) {
	wg := NewWaitGroup()
	wg.Add(2)

	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	wg.Done()
	select {
	case <-done:
		t.Fatal("Wait returned too early")
	case <-time.After(20 * time.Millisecond):
	}

	wg.Done()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Wait did not return")
	}
}

func TestWaitGroupWithWorkers(t *testing.T) {
	wg := NewWaitGroup()
	const workers = 10
	wg.Add(workers)

	completed := make(chan struct{}, workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			completed <- struct{}{}
		}()
	}

	wg.Wait()
	if len(completed) != workers {
		t.Fatalf("completed %d workers, want %d", len(completed), workers)
	}
}
