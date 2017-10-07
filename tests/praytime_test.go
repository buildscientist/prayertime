package test

import (
	"github.com/buildscientist/prayertime/praytime"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var isnaPrayerCalcMethodTests = []struct {
	latitude, longitude, timezone float64  //input
	expected                      []string //expected result
}{
	//San Diego,California
	{32.715738, -117.161084, -7, []string{"05:39 AM", "06:46 AM", "12:37 PM", "03:56 PM", "06:26 PM", "06:26 PM", "07:34 PM"}},
	//Denver,Colorado
	{39.739236, -104.990251, -6, []string{"05:48 AM", "07:01 AM", "12:48 PM", "04:02 PM", "06:34 PM", "06:34 PM", "07:48 PM"}},
	//Chicago, Illinois
	{41.878114, -87.629798, -5, []string{"05:37 AM", "06:53 AM", "12:38 PM", "03:51 PM", "06:23 PM", "06:23 PM", "07:39 PM"}},
	//Toronto, Ontario
	{43.653226, -79.383184, -4, []string{"06:03 AM", "07:21 AM", "01:06 PM", "04:16 PM", "06:49 PM", "06:49 PM", "08:08 PM"}},
}

var mwlPrayerCalcMethodTests = []struct {
	latitude, longitude, timezone float64  //input
	expected                      []string //expected result
}{
	//London, England
	{51.507351, -0.127758, 1, []string{"05:18 AM", "07:10 AM", "12:49 PM", "03:48 PM", "06:27 PM", "06:27 PM", "08:11 PM"}},
	//Sydney, Australia
	{-33.868820, 151.209296, 11, []string{"05:01 AM", "06:26 AM", "12:43 PM", "04:17 PM", "07:01 PM", "07:01 PM", "08:21 PM"}},
}

var oneSeventhHighLatTests = []struct {
	latitude, longitude, timezone float64  //input
	expected                      []string //expected result
}{
	//Nuuk,Greenland
	{64.1814, -51.6941, -2, []string{"05:57 AM", "07:51 AM", "01:15 PM", "03:43 PM", "06:37 PM", "06:37 PM", "08:30 PM"}},
	//Helsinki, Finland
	{60.192059, 24.945831, 3, []string{"05:46 AM", "07:38 AM", "01:08 PM", "03:48 PM", "06:37 PM", "06:37 PM", "08:29 PM"}},
	//Tromsø, Norway
	{69.6492, 18.9553, 2, []string{"05:23 AM", "07:19 AM", "12:32 PM", "02:40 PM", "05:43 PM", "05:43 PM", "07:40 PM"}},
	//Khatanga, Russia
	{71.9640, 102.4406, 7, []string{"04:53 AM", "06:51 AM", "11:58 AM", "01:56 PM", "05:03 PM", "05:03 PM", "07:02 PM"}},
}

func TestAdjustHighAltitude(t *testing.T) {
	assert := assert.New(t)
	today := time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local)

	for _, tt := range oneSeventhHighLatTests {
		city := praytime.New(tt.latitude, tt.longitude, tt.timezone)
		city.PrayerCalcMethod = praytime.MWL
		city.AdjustHighLats = praytime.ONE_SEVENTH
		cityPrayerTime := praytime.CalculatePrayerTimes(&city, today)
		assert.NotNil(cityPrayerTime)
		assert.NotEmpty(cityPrayerTime)
		assert.NotContains(cityPrayerTime, praytime.INVALID_TIME)
		assert.EqualValues(tt.expected, cityPrayerTime)
	}
}

func TestISNAPrayerMethod(t *testing.T) {
	assert := assert.New(t)
	today := time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local)

	for _, tt := range isnaPrayerCalcMethodTests {
		city := praytime.New(tt.latitude, tt.longitude, tt.timezone)
		cityPrayerTime := praytime.CalculatePrayerTimes(&city, today)
		assert.NotNil(cityPrayerTime)
		assert.NotEmpty(cityPrayerTime)
		assert.NotContains(cityPrayerTime, praytime.INVALID_TIME)
		assert.EqualValues(tt.expected, cityPrayerTime)
	}

}

func TestMWLPrayerMethod(t *testing.T) {
	assert := assert.New(t)
	today := time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local)

	for _, tt := range mwlPrayerCalcMethodTests {
		city := praytime.New(tt.latitude, tt.longitude, tt.timezone)
		city.PrayerCalcMethod = praytime.MWL
		cityPrayerTime := praytime.CalculatePrayerTimes(&city, today)
		assert.NotNil(cityPrayerTime)
		assert.NotEmpty(cityPrayerTime)
		assert.NotContains(cityPrayerTime, praytime.INVALID_TIME)
		assert.EqualValues(tt.expected, cityPrayerTime)
	}

}

func TestMakkahPrayerMethod(t *testing.T) {
	assert := assert.New(t)

	today := time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local)
	expectedPrayerTime := []string{"05:00 AM", "06:16 AM", "12:11 PM", "03:34 PM", "06:06 PM", "06:06 PM", "07:36 PM"}

	jeddah := praytime.New(21.285407, 39.237551, 3)
	jeddah.PrayerCalcMethod = praytime.MAKKAH
	jeddahPrayerTime := praytime.CalculatePrayerTimes(&jeddah, today)

	assert.NotNil(jeddahPrayerTime)
	assert.NotEmpty(jeddahPrayerTime)
	assert.NotContains(jeddahPrayerTime, praytime.INVALID_TIME)
	assert.EqualValues(expectedPrayerTime, jeddahPrayerTime)

}

func TestKarachiPrayerMethod(t *testing.T) {
	assert := assert.New(t)

	today := time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local)
	expectedPrayerTime := []string{"05:10 AM", "06:26 AM", "12:20 PM", "04:35 PM", "06:14 PM", "06:14 PM", "07:30 PM"}

	karachi := praytime.New(24.861462, 67.009939, 5)
	karachi.PrayerCalcMethod = praytime.KARACHI
	karachi.AsrJuristic = praytime.HANAFI
	karachiPrayerTime := praytime.CalculatePrayerTimes(&karachi, today)

	assert.NotNil(karachiPrayerTime)
	assert.NotEmpty(karachiPrayerTime)
	assert.NotContains(karachiPrayerTime, praytime.INVALID_TIME)
	assert.EqualValues(expectedPrayerTime, karachiPrayerTime)

}

func TestEgyptPrayerMethod(t *testing.T) {
	assert := assert.New(t)

	today := time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local)
	expectedPrayerTime := []string{"04:25 AM", "05:51 AM", "11:43 AM", "03:04 PM", "05:34 PM", "05:34 PM", "06:52 PM"}

	cairo := praytime.New(30.044420, 31.235712, 2)
	cairo.PrayerCalcMethod = praytime.EGYPT
	cairoPrayerTime := praytime.CalculatePrayerTimes(&cairo, today)

	assert.NotNil(cairoPrayerTime)
	assert.NotEmpty(cairoPrayerTime)
	assert.NotContains(cairoPrayerTime, praytime.INVALID_TIME)
	assert.EqualValues(expectedPrayerTime, cairoPrayerTime)

}

func TestNoSuffix(t *testing.T) {
	assert := assert.New(t)

	today := time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local)

	city := praytime.New(30.044420, 31.235712, 2)
	city.TimeFormat = praytime.TIME_12_NO_SUFFIX
	cityPrayerTime := praytime.CalculatePrayerTimes(&city, today)

	assert.NotNil(cityPrayerTime)
	assert.NotEmpty(cityPrayerTime)
	assert.NotContains(cityPrayerTime, praytime.INVALID_TIME)
	assert.NotContains(cityPrayerTime, "AM")
	assert.NotContains(cityPrayerTime, "PM")
}

func Test24Hour(t *testing.T) {
	assert := assert.New(t)

	today := time.Date(2017, 10, 6, 0, 0, 0, 0, time.Local)

	city := praytime.New(30.044420, 31.235712, 2)
	city.TimeFormat = praytime.TIME_24
	cityPrayerTime := praytime.CalculatePrayerTimes(&city, today)

	assert.NotNil(cityPrayerTime)
	assert.NotEmpty(cityPrayerTime)
	assert.NotContains(cityPrayerTime, praytime.INVALID_TIME)
	assert.NotContains(cityPrayerTime, "AM")
	assert.NotContains(cityPrayerTime, "PM")

}
