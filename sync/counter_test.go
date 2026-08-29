package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("counter to 3", func(t *testing.T) {
		counter := NewDefaultCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("should run safely concurrently", func(t *testing.T) {
		wantedCount := 1000

		counter := NewDefaultCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for range wantedCount {
			go func() {
				counter.Inc()
				defer wg.Done()
			}()
		}

		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})

	t.Run("atomic counter", func(t *testing.T) {
		wantedCount := 1000

		atomicCounter := NewAtomicCounter()

		for range wantedCount {
			atomicCounter.Inc()
		}

		assertCounter(t, atomicCounter, wantedCount)
	})
}

func assertCounter(t testing.TB, counter Counter, want int) {
	t.Helper()

	got := counter.Value()

	if got != int64(want) {
		t.Errorf("got %d, want %d", got, want)
	}
}
