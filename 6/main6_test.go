package main

import (
	"context"
	"testing"
	"time"
)

func TestRandomGenerator(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	out := randomGenerator(ctx, 1, 10)
	for i := 0; i < 10; i++ {
		select {
		case value, ok := <-out:
			if !ok {
				t.Fatal("generator closed too early")
			}
			if value < 1 || value > 10 {
				t.Fatalf("value %d is outside [1, 10]", value)
			}
		case <-time.After(time.Second):
			t.Fatal("timed out waiting for value")
		}
	}
}

func TestRandomGeneratorStops(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	out := randomGenerator(ctx, 1, 10)
	cancel()

	select {
	case _, ok := <-out:
		if ok {
			// A value may already be in flight; wait for closure below.
			select {
			case _, ok = <-out:
				if ok {
					t.Fatal("generator did not stop")
				}
			case <-time.After(time.Second):
				t.Fatal("timed out waiting for generator to stop")
			}
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for generator to stop")
	}
}
