package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hourInHalfClock    = 6
	hoursInClock       = 2 * hourInHalfClock
)

// Point is a point or hand on the clockface
type Point struct {
	X, Y float64
}

// SecondHandPoint creates a Point given a time
func SecondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / (float64(t.Second()))))
}

// MinuteHandPoint creates a Point given a time
func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

// HourHandPoint creates a Point given a time
func HourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	return ((minutesInRadians(t) / hoursInClock) + (math.Pi / (hourInHalfClock / float64(t.Hour()%12))))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
