package sync

import "testing"

func TestCounter(t *testing.T) {
	t.Run("increment counter 3 times and get 3", func(t *testing.T) {
		expected := 3
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()
		got := counter.Value()

		if got != expected {
			t.Errorf("Expected counter value %d, but got %d", expected, got)
		}
	})

}
