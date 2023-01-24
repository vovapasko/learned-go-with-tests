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
	t.Run("test rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		got := rectangle.Area()
		want := 100.0
		if got != want {
			t.Errorf("Got %g, want %g", got, want)
		}
	})
	t.Run("test circle area", func(t *testing.T) {
		circle := Circle{100}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("Got %g, want %g", got, want)
		}
	})
}
