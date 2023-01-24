package structs_methods_interfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 10}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Got %g, want %g", got, want)
		}

	}

	t.Run("test rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		checkArea(t, rectangle, 100.0)
	})
	t.Run("test circle area", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
