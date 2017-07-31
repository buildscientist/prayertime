package praytime

import "github.com/buildscientist/prayertime/trig"

type PrayTime struct {
	latitude, longitude, timezone float64
}

func init() {
	//Prayer Time names
	//var prayerTimeNames = []string{"Fajr","Sunrise","Dhuhr","Asr","Sunset","Maghrib","Isha"}

	//Prayer Time Method Parameters
	var methodParams = make(map[int][]float64)
	methodParams[JAFARI] = []float64{16, 0, 4, 0, 14}
	methodParams[KARACHI] = []float64{18, 1, 0, 0, 18}
	methodParams[ISNA] = []float64{15, 1, 0, 0, 15}
	methodParams[MWL] = []float64{18, 1, 0, 0, 17}
	methodParams[MAKKAH] = []float64{18.5, 1, 0, 1, 90}
	methodParams[EGYPT] = []float64{19.5, 1, 0, 0, 17.5}
	methodParams[TEHRAN] = []float64{17.7, 0, 4.5, 0, 14}
	methodParams[CUSTOM] = []float64{18, 1, 0, 0, 17}

}

func New(latitude, longitude, timezone float64) PrayTime {
	return PrayTime{latitude, longitude, timezone}
}

//Prayer Time Calculation functions
func sunPosition(julianDate float64) (position []float64) {
	var daysFromJulianEpoch = julianDate - 2451545.0
	var meanSunAnomaly = trig.FixAngle(357.529 + (0.98560028 * daysFromJulianEpoch))
	var meanSunLongitude = trig.FixAngle(280.459 + (0.98564736 * daysFromJulianEpoch))
	var geoCentricElipticSunLongitude = trig.FixAngle(meanSunLongitude + (1.915 * trig.DegreeSin(meanSunAnomaly)) + (0.020 * trig.DegreeSin(2*meanSunAnomaly)))

	var meanObliquityEcliptic = 23.439 - (0.00000036 * daysFromJulianEpoch)
	var sunDeclination = trig.DegreeArcSin(trig.DegreeSin(meanObliquityEcliptic) * trig.DegreeSin(geoCentricElipticSunLongitude))
	var rightAscension = (trig.DegreeArcTan2(trig.DegreeCos(meanObliquityEcliptic) * trig.DegreeSin(geoCentricElipticSunLongitude), trig.DegreeCos(geoCentricElipticSunLongitude))) / 15.0

	rightAscension = trig.FixHour(rightAscension)
	var equationOfTime = meanSunLongitude/15.0 - rightAscension
	
	return []float64{sunDeclination,equationOfTime}
}