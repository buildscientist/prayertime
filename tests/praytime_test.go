package test

import (
	"github.com/buildscientist/prayertime/praytime"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCalculatePrayerTimes(t *testing.T) {
	assert := assert.New(t)

	today := time.Now()

	chicago := praytime.New(41.879626, -87.648217, -5)
	chicago.PrayerCalcMethod = praytime.MAKKAH
	chicagoPrayerTime := praytime.CalculatePrayerTimes(&chicago, today)

	assert.NotNil(chicagoPrayerTime)
	assert.NotContains(chicagoPrayerTime, praytime.INVALID_TIME)

}

func TestAdjustHighAltitude(t *testing.T) {
	assert := assert.New(t)
	today := time.Date(2017, 9, 21, 0, 0, 0, 0, time.Local)

	helsinki := praytime.New(60.192059, 24.945831, 3)
	helsinki.TimeFormat = praytime.TIME_24
	helsinki.PrayerCalcMethod = praytime.MWL
	helsinki.AdjustHighLats = praytime.ONE_SEVENTH
	helsinki.AsrJuristic = praytime.SHAFII

	helsinkiPrayerTime := praytime.CalculatePrayerTimes(&helsinki, today)

	assert.NotNil(helsinkiPrayerTime)
	assert.NotContains(helsinkiPrayerTime, praytime.INVALID_TIME)

}
