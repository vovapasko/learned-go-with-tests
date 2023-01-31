package clockface

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(time time.Time) Point {
	return Point{150, 150 - 90}
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}
