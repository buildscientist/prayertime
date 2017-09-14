package praytime

import (
	"github.com/buildscientist/prayertime/julian"
	"github.com/buildscientist/prayertime/trig"
	"math"
	"strconv"
	"time"
)

var methodParams = make(map[int][]float64)
var PrayerCalcMethod, AsrJuristic, AdjustHighLats, TimeFormat int
var PrayerTimeNames = []string{FAJR, "Sunrise", DHUHR, ASR, "Sunset", MAGHRIB, ISHA}
var julianDate float64
var prayerTimesCurrent []float64
var Offsets = [7]int{0, 0, 0, 0, 0, 0, 0}

func init() {
	//Prayer Time Method Parameters
	methodParams[JAFARI] = []float64{16, 0, 4, 0, 14}
	methodParams[KARACHI] = []float64{18, 1, 0, 0, 18}
	methodParams[ISNA] = []float64{15, 1, 0, 0, 15}
	methodParams[MWL] = []float64{18, 1, 0, 0, 17}
	methodParams[MAKKAH] = []float64{18.5, 1, 0, 1, 90}
	methodParams[EGYPT] = []float64{19.5, 1, 0, 0, 17.5}
	methodParams[TEHRAN] = []float64{17.7, 0, 4.5, 0, 14}
	methodParams[CUSTOM] = []float64{18, 1, 0, 0, 17}

	PrayerCalcMethod = ISNA
	AsrJuristic = SHAFII
	AdjustHighLats = NONE
	TimeFormat = TIME_24
}

type PrayerLocale struct {
	latitude, longitude, timezone float64
}

func New(latitude, longitude, timezone float64) PrayerLocale {
	return PrayerLocale{latitude, longitude, timezone}
}

//Prayer Time Calculation functions
func sunPosition(julianDate float64) (position []float64) {
	var daysFromJulianEpoch = julianDate - 2451545.0
	var meanSunAnomaly = trig.FixAngle(357.529 + (0.98560028 * daysFromJulianEpoch))
	var meanSunLongitude = trig.FixAngle(280.459 + (0.98564736 * daysFromJulianEpoch))
	var geoCentricElipticSunLongitude = trig.FixAngle(meanSunLongitude + (1.915 * trig.DegreeSin(meanSunAnomaly)) + (0.020 * trig.DegreeSin(2*meanSunAnomaly)))

	var meanObliquityEcliptic = 23.439 - (0.00000036 * daysFromJulianEpoch)
	var sunDeclination = trig.DegreeArcSin(trig.DegreeSin(meanObliquityEcliptic) * trig.DegreeSin(geoCentricElipticSunLongitude))
	var rightAscension = (trig.DegreeArcTan2(trig.DegreeCos(meanObliquityEcliptic)*trig.DegreeSin(geoCentricElipticSunLongitude), trig.DegreeCos(geoCentricElipticSunLongitude))) / 15.0

	rightAscension = trig.FixHour(rightAscension)
	var equationOfTime = meanSunLongitude/15.0 - rightAscension

	return []float64{sunDeclination, equationOfTime}
}

func equationOfTime(julianDate float64) float64 {
	var equationOfTime = sunPosition(julianDate)[1]
	return equationOfTime
}

func sunDeclination(julianDate float64) float64 {
	var declinationAngle = sunPosition(julianDate)[0]
	return declinationAngle
}

func computeMidDay(time float64) float64 {
	var currentTime = equationOfTime(julianDate + time)
	return trig.FixHour(12 - currentTime)
}

func computeTime(prayTime *PrayerLocale, angle, time float64) float64 {
	var D = sunDeclination(julianDate) + time
	var Z = computeMidDay(time)
	var beg = -trig.DegreeSin(angle) - trig.DegreeSin(D)*trig.DegreeSin(prayTime.latitude)
	var mid = trig.DegreeCos(D) * trig.DegreeCos(prayTime.latitude)
	var v = trig.DegreeArcCos(beg/mid) / 15.0

	if angle > 90 {
		return Z - v
	}
	return Z + v
}

func computeAsr(prayTime *PrayerLocale, step, time float64) float64 {
	var D = sunDeclination(julianDate + time)
	var G = -trig.DegreeArcCot(step + trig.DegreeTan(math.Abs(prayTime.latitude-D)))
	return computeTime(prayTime, G, time)
}

func timeDifference(timeOne, timeTwo float64) float64 {
	return trig.FixHour(timeTwo - timeOne)
}

func getDatePrayerTimes(prayTime *PrayerLocale, year, month, day int) []string {
	julianDate = julian.ConvertFromGregToJul(year, month, day)
	var longitudinalDiff = prayTime.longitude / (15.0 * 24.0)
	julianDate = julianDate - longitudinalDiff
	return computeDayTimes(prayTime)
}

func CalculatePrayerTimes(prayTime *PrayerLocale, today time.Time) []string {
	var year = today.Year()
	var month = int(today.Month())
	var day = today.Day()
	return getDatePrayerTimes(prayTime, year, month, day)
}

func setCustomParams(params []float64) {
	for x := 0; x < 5; x++ {
		if params[x] == -1 {
			params[x] = methodParams[PrayerCalcMethod][x]
			methodParams[CUSTOM] = params
		} else {
			methodParams[CUSTOM][x] = params[x]
		}

	}
	PrayerCalcMethod = CUSTOM
}

func setPrayerAngle(prayerName string, angle float64) {
	switch {
	case prayerName == FAJR:
		setCustomParams([]float64{angle, -1, -1, -1, -1})

	case prayerName == MAGHRIB:
		setCustomParams([]float64{-1, 0, angle, -1, -1})

	case prayerName == ISHA:
		setCustomParams([]float64{-1, -1, -1, 0, angle})
	}

}

func setPrayerMinutes(prayerName string, minutes float64) {
	switch {
	case prayerName == MAGHRIB:
		setCustomParams([]float64{-1, 1, minutes, -1, -1})

	case prayerName == ISHA:
		setCustomParams([]float64{-1, -1, -1, 1, minutes})
	}
}

func floatToTime(time float64, useSuffix, twentyFourHourFormat bool) string {
	if math.IsNaN(time) {
		return INVALID_TIME
	}

	var result, suffix string

	time = trig.FixHour(time + 0.5/60.0)
	var hours = int(math.Floor(time))
	var minutes = math.Floor((time - float64(hours)) * 60.0)

	if hours >= 12 {
		suffix = "PM"
	} else {
		suffix = "AM"
	}

	if twentyFourHourFormat {
		hours = ((((hours + 12) - 1) % (12)) + 1)
	}

	if useSuffix {
		switch {
		case (hours >= 0 && hours <= 9) && (minutes >= 0 && minutes <= 9):
			result = "0" + strconv.Itoa(hours) + ":0" + strconv.Itoa(int(minutes)) + " " + suffix

		case (hours >= 0 && hours <= 9):
			result = "0" + strconv.Itoa(hours) + ":" + strconv.Itoa(int(minutes)) + " " + suffix

		case (minutes >= 0 && minutes <= 9):
			result = strconv.Itoa(hours) + ":0" + strconv.Itoa(int(minutes)) + " " + suffix

		default:
			result = strconv.Itoa(hours) + ":" + strconv.Itoa(int(minutes)) + " " + suffix
		}

	} else {
		switch {
		case (hours >= 0 && hours <= 9) && (minutes >= 0 && minutes <= 9):
			result = "0" + strconv.Itoa(hours) + ":0" + strconv.Itoa(int(minutes))

		case (hours >= 0 && hours <= 9):
			result = "0" + strconv.Itoa(hours) + ":" + strconv.Itoa(int(minutes))

		case (minutes >= 0 && minutes <= 9):
			result = strconv.Itoa(hours) + ":0" + strconv.Itoa(int(minutes))

		default:
			result = strconv.Itoa(hours) + ":" + strconv.Itoa(int(minutes))
		}
	}

	return result
}

func dayPortion(times []float64) []float64 {
	for x := 0; x < 7; x++ {
		times[x] /= 24
	}
	return times
}

func computePrayerTime(prayTime *PrayerLocale, times []float64) []float64 {
	var time = dayPortion(times)
	var angle = 180 - methodParams[PrayerCalcMethod][0]
	var fajr = computeTime(prayTime, angle, time[0])
	var sunrise = computeTime(prayTime, 180-0.833, time[1])
	var dhuhr = computeMidDay(time[2])
	var asr = computeAsr(prayTime, float64(1+AsrJuristic), time[3])
	var sunset = computeTime(prayTime, 0.833, time[4])
	var maghrib = computeTime(prayTime, methodParams[PrayerCalcMethod][2], time[5])
	var isha = computeTime(prayTime, methodParams[PrayerCalcMethod][4], time[6])

	var computedPrayerTimes = []float64{fajr, sunrise, dhuhr, asr, sunset, maghrib, isha}

	return computedPrayerTimes
}

func adjustTimes(prayTime *PrayerLocale, times []float64) []float64 {
	for x := 0; x < len(times); x++ {
		times[x] = times[x] + (prayTime.timezone - (prayTime.longitude/15))
	}

	times[2] = times[2] + float64(DHUHR_MINUTES / 60)

	switch {
	case methodParams[PrayerCalcMethod][1] == 1:
		times[5] = times[4] + methodParams[PrayerCalcMethod][2]/60

	case methodParams[PrayerCalcMethod][3] == 1:
		times[6] = times[5] + methodParams[PrayerCalcMethod][4]/60

	case AdjustHighLats != 0:
		times = adjustHighLatTimes(times)
	}

	return times
}

// Adjust Fajr, Isha and Maghrib for locations in higher latitudes
func adjustHighLatTimes(times []float64) []float64 {
	var nightTime = timeDifference(times[4], times[1])
	var fajrDiff = nightPortion(methodParams[PrayerCalcMethod][0] * nightTime)

	if math.IsNaN(times[0]) || timeDifference(times[0], times[1]) > fajrDiff {
		times[0] = times[1] - fajrDiff
	}

	var ishaAngle float64
	if methodParams[PrayerCalcMethod][3] == 0 {
		ishaAngle = methodParams[PrayerCalcMethod][4]
	} else {
		ishaAngle = 18.0
	}
	var ishaDiff = nightPortion(ishaAngle) * nightTime

	if math.IsNaN(times[6]) || timeDifference(times[4], times[6]) > ishaDiff {
		times[6] = times[4] + ishaDiff
	}

	var maghribAngle float64
	if methodParams[PrayerCalcMethod][1] == 0 {
		maghribAngle = methodParams[PrayerCalcMethod][2]
	} else {
		maghribAngle = 4.0
	}
	var maghribDiff = nightPortion(maghribAngle) * nightTime

	if math.IsNaN(times[5]) || timeDifference(times[4], times[5]) > maghribDiff {
		times[5] = times[4] + maghribDiff
	}

	return times
}

func nightPortion(angle float64) float64 {
	var calc = 0.0
	switch {
	case AdjustHighLats == ANGLE_BASED:
		calc = angle / 60.0
	case AdjustHighLats == MIDNIGHT:
		calc = 0.5

	case AdjustHighLats == ONE_SEVENTH:
		calc = 0.14286
	}
	return calc
}

func tune(offsetTimes []int) {
	for x := 0; x < len(offsetTimes); x++ {
		Offsets[x] = offsetTimes[x]
	}
}

func tuneTimes(times []float64) []float64 {
	for x := 0; x < len(times); x++ {
		times[x] = times[x] + float64(Offsets[x]/60.0)
	}
	return times
}

func adjustTimesFormat(times []float64) []string {
	var result []string
	if TimeFormat == 3 {
		for index := range times {
			result = append(result, strconv.FormatFloat(times[index], 'f', -1, 64))
		}
		return result
	}

	for x := 0; x < 7; x++ {
		switch {
		case TimeFormat == TIME_12:
			result = append(result, floatToTime(times[x], true, false))

		case TimeFormat == TIME_12_NO_SUFFIX:
			result = append(result, floatToTime(times[x], false, false))

		default:
			result = append(result, floatToTime(times[x], false, true))
		}
	}
	return result
}

func computeDayTimes(prayTime *PrayerLocale) []string {
	var times = []float64{5, 6, 12, 13, 18, 18, 18}

	for x := 1; x <= NUMBER_OF_ITERATIONS; x++ {
		times = computePrayerTime(prayTime, times)
	}

	times = adjustTimes(prayTime, times)
	times = tuneTimes(times)

	return adjustTimesFormat(times)
}
