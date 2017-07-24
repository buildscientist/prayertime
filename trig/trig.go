package trig

import "math"

//Trigonometric Helpers

func DegreeSin(degrees float64) (sinDegrees float64) {
	return math.Sin(DegreesToRadians(degrees))
}

func DegreeCos(degrees float64) (cosDegrees float64) {
	return math.Cos(DegreesToRadians(degrees))
}

func DegreeTan(degrees float64) (tanDegrees float64) {
	return math.Tan(DegreesToRadians(degrees))
}

func DegreeArcSin(degrees float64) (arcSinDegrees float64) {
	return RadiansToDegrees(math.Asin(degrees))
}

func DegreeArcCos(degrees float64) (arcCosDegrees float64) {
	return RadiansToDegrees(math.Acos(degrees))
}

func DegreeArcTan(degrees float64) (arcTanDegrees float64) {
	return RadiansToDegrees(math.Atan(degrees))
}

func DegreeArcTan2(degreeA, degreeB float64) (arcTan2Degrees float64) {
	return RadiansToDegrees(math.Atan2(degreeA, degreeB))
}

func DegreeArcCot(degrees float64) (arcCotDegrees float64) {
	return RadiansToDegrees(math.Atan2(1.0, degrees))
}

func RadiansToDegrees(radians float64) (degrees float64) {
	return (radians * 180) / math.Pi
}

func DegreesToRadians(degrees float64) (radians float64) {
	return (degrees * math.Pi) / 180
}

func Fix(value, mode float64) (a float64) {
	value = value - mode*(math.Floor(value/mode))

	if value < 0 {
		return value + mode
	}
	return value
}

func FixAngle(angle float64) (rangeReducedAngle float64) {
	return Fix(angle, 360.0)
}

func FixHour(hour float64) (fixedHour float64) {
	return Fix(hour, 24.0)
}
