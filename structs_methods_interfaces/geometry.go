package structs_methods_interfaces

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (circle Circle) Area() float64 {
	return circle.Radius * circle.Radius * math.Pi
}
