package julian

import "math"

/*
	Converts a Gregorian calendar date to a Julian date. More details about the exact formula are described in detail
	in Vallado's "Fundamentals of Astrodynamics and Applications"
*/

func ConvertFromGregToJul(year, month, day int) float64 {
	if month <= 2 {
		year -= 1
		month += 12
	}
	var a = math.Floor(float64(year / 100.0))

	var b = 2 - a + math.Floor(a/4.0)

	var julianDate = math.Floor(AVERAGE_GREGORIAN_YEAR_LENGTH*float64((year+JULIAN_EPOCH))) + math.Floor(AVERAGE_GREGORIAN_MONTH_LENGTH*float64((month+1))) + float64(day) + b - JULIAN_OFFSET

	return julianDate
}