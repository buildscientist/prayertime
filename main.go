package main

import "math"
import "fmt"

func main() {
	fmt.Println(math.Sin(degreesToRadians(90)))
}

//Trigonometric Helpers

func degreeSin(degrees float64) (sinDegrees float64) {
	return math.Sin(degreesToRadians(degrees))
}

func degreeCos(degrees float64) (cosDegrees float64) {
	return math.Cos(degreesToRadians(degrees))
}

func degreeTan(degrees float64) (tanDegrees float64) {
	return math.Tan(degreesToRadians(degrees))
}

func degreeArcSin(degrees float64) (arcSinDegrees float64) {
	return radiansToDegrees(math.Asin(degrees))
}

func degreeArcCos(degrees float64) (arcCosDegrees float64) {
	return radiansToDegrees(math.Acos(degrees))
}

func degreeArcTan(degrees float64) (arcTanDegrees float64) {
	return radiansToDegrees(math.Atan(degrees))
}

func degreeArcTan2(degreeA,degreeB float64) (arcTan2Degrees float64) {
	return radiansToDegrees(math.Atan2(degreeA,degreeB))
}

func degreeArcCot(degrees float64) (arcCotDegrees float64) {
	return radiansToDegrees(math.Atan2(1.0,degrees))
}


func radiansToDegrees(radians float64) (degrees float64) {
	return (radians * 180) / math.Pi
}

func degreesToRadians(degrees float64) (radians float64) {
	return (degrees * math.Pi) / 180
}

func fix(value, mode float64) (a float64) {
	value = value - mode*(math.Floor(value/mode))

	if value < 0 {
		return value + mode
	}
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
