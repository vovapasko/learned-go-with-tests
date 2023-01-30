package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("increment counter 3 times and get 3", func(t *testing.T) {
		expected := 3
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, expected)
	})
	t.Run("increment counter in different threads", func(t *testing.T) {
		wg := sync.WaitGroup{}
		expectedCount := 1000
		counter := NewCounter()
		wg.Add(expectedCount)
		for i := 0; i < expectedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()
		assertCounter(t, counter, expectedCount)
	})

}

func BenchmarkCounter_Inc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := Counter{}
		counter.Inc()
	}
}

func BenchmarkCounter_Inc_Concurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := sync.WaitGroup{}
		counter := Counter{}
		wg.Add(1)
		go func() {
			counter.Inc()
			wg.Done()
		}()

	}
}

func assertCounter(t *testing.T, counter *Counter, expected int) {
	got := counter.Value()

	if got != expected {
		t.Errorf("Expected counter value %d, but got %d", expected, got)
	}
}
