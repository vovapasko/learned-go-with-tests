package iteration

import "testing"

func TestRepeat(t *testing.T) {
	actual := Repeat("a")
	expected := "aaaaaa"

	if actual != expected {
		t.Errorf("expected '%q', but found '%q'", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
