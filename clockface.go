package clockface

import (
	"math"
	"time"
)

// Point is a point or hand on the clockface
type Point struct {
	X, Y float64
}

// SecondHand create a second hand poin
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
