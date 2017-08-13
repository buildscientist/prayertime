package praytime

import(
	"math"
	"time"
	"strconv"
	"github.com/buildscientist/prayertime/trig"
	"github.com/buildscientist/prayertime/julian"
)

type PrayTime struct {
	latitude, longitude, timezone float64
}

var methodParams = make(map[int][]float64)
var prayerCalcMethod, asrJuristic, adjustHighLats, timeFormat int
var prayerTimeNames = []string{FAJR, "Sunrise", DHUHR, ASR, "Sunset", MAGHRIB, ISHA}
var julianDate float64
var invalidTime string
var prayerTimesCurrent []float64
var offSets = [7]int{0, 0, 0, 0, 0, 0, 0}

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

	prayerCalcMethod = 0
	asrJuristic = 0
	adjustHighLats = 1
	timeFormat = 0
	invalidTime = "-----"
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

func computeTime(prayTime *PrayTime, angle, time float64) float64 {
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

func computeAsr(prayTime *PrayTime, step, time float64) float64 {
	var D = sunDeclination(julianDate + time)
	var G = -trig.DegreeArcCot(step + trig.DegreeTan(math.Abs(prayTime.latitude-D)))
	return computeTime(prayTime, G, time)
}

func timeDifference(timeOne, timeTwo float64) float64 {
	return trig.FixHour(timeTwo - timeOne)
}

func getDatePrayerTimes(prayTime *PrayTime, year, month, day int) []string {
	julianDate = julian.ConvertFromGregToJul(year, month, day)
	var longitudinalDiff = prayTime.longitude / (15.0 * 24.0)
	julianDate = julianDate - longitudinalDiff

	//TODO: return computeDayTimes()
	return []string{"1"}
}

func getPrayerTimes(prayTime *PrayTime, today time.Time) []string {
	var year = today.Year()
	var month = int(today.Month())
	var day = today.Day()
	return getDatePrayerTimes(prayTime, year, month, day)
}

func setCustomParams(params []float64) {
	for x := 0; x < 5; x++ {
		if(params[x] ==  -1) {
			params[x] = methodParams[prayerCalcMethod][x]
			methodParams[CUSTOM] = params
		} else {
			methodParams[CUSTOM][x] = params[x]
		}

	}
	prayerCalcMethod = CUSTOM
}


/* func setPrayerAngle(prayerName string, angle float64) {
	switch {
		case prayerName == FAJR:
			setCustomParams([]float64{angle, -1, -1, -1, -1})
		
		case prayerName == MAGHRIB:
			setCustomParams([]float64{-1, 0, angle, -1, -1})

		case prayerName == ISHA:
			setCustomParams([]float64{-1, -1, -1, 0, angle})
	}

} */


func setFajrAngle(angle float64) {
	var params = []float64{angle, -1, -1, -1, -1}
	setCustomParams(params)
}

func setMaghribAngle(angle float64) {
	var params = []float64{-1, 0, angle, -1, -1}
	setCustomParams(params)
}

func setIshaAngle(angle float64) {
	var params = []float64{-1, -1, -1, 0, angle}
	setCustomParams(params)
}

func setMaghribMinute(minutes float64) {
	var params = []float64{-1, 1, minutes, -1, -1}
	setCustomParams(params)
}

func setIshaMinutes(minutes float64) {
	var params = []float64{-1, -1, -1, 1, minutes}
	setCustomParams(params)
}

func floatToTime24(time float64) string {
	var result string

	if(math.IsNaN(time)) {
		return invalidTime
	}

	time = trig.FixHour(time + 0.5 / 60.0)
	var hours = int(math.Floor(time))
	var minutes = math.Floor((time - float64(hours)) * 60.0)

	switch {
		case (hours >= 0 && hours <= 9) && (minutes >= 0  &&  minutes <= 9):
			result = "0" + strconv.Itoa(hours) + ":0" + strconv.Itoa(int(minutes))
		
		case (hours >= 0 && hours <= 9):
			result = "0" + strconv.Itoa(hours) + ":" + strconv.Itoa(int(minutes))

		case (minutes >= 0 && minutes <= 9):
			result = strconv.Itoa(hours) + ":0" + strconv.Itoa(int(minutes))
		
		default:
			result = strconv.Itoa(hours) + ":" + strconv.Itoa(int(minutes))
	}

	return result
}