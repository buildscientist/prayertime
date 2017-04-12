package main

import "math"
import "fmt"

func main() {
	var x = 90 * (math.Pi / 180)
	fmt.Println(math.Sin(x))
}

//Trigonometric Helpers
func fix(value, mode float64) (a float64) {
	value = value - mode*(math.Floor(value/mode))
	return value
}

func fixAngle(angle float64) (rangeReducedAngle float64) {
	return fix(angle, 360.0)
}

func fixHour(hour float64) (fixedHour float64) {
	return fix(hour, 24.0)
}

//Prayer Time Calculation functions
func sunPosition(julianDate float64) (position []float64) {

	var daysFromEpoch = julianDate - 2451545.0

	return []float64{daysFromEpoch, 2.0, 3.0}
}
