package clockface

import (
	"math"
	"time"
)

const secondHandLength = 90
const clockCenterX = 150
const clockCenterY = 150

// Point is a point or hand on the clockface
type Point struct {
	X, Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time 't' represented as a Point
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
	p = Point{p.X, -p.Y}                                      // flip
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}         // translate
	return p
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
