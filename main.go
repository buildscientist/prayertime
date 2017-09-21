/*
Copyright 2017
Author: Youssuf ElKalay
All rights reserved. Use of this source code is governed by the Apache 2.0 license
that can be found in the LICENSE file or at http://www.apache.org/licenses/LICENSE-2.0
*/

package main

import (
	"fmt"
	"github.com/buildscientist/prayertime/praytime"
	"time"
)

func main() {
	var helsinki = praytime.New(60.192059, 24.945831, 3)
	var toronto = praytime.New(43.6532, -79.3832, -4.0)
	var chicago = praytime.New(41.879626, -87.648217, -5.0)
	var sandiego = praytime.New(32.7157, -117.1611, -7.0)

	helsinki.TimeFormat = praytime.TIME_24
	helsinki.PrayerCalcMethod = praytime.MWL
	helsinki.AdjustHighLats = praytime.ANGLE_BASED
	helsinki.AsrJuristic = praytime.HANAFI

	toronto.TimeFormat = praytime.TIME_24
	toronto.AdjustHighLats = praytime.ANGLE_BASED
	toronto.AsrJuristic = praytime.HANAFI

	chicago.TimeFormat = praytime.TIME_12_NO_SUFFIX
	sandiego.TimeFormat = praytime.TIME_12

	var helsinkiPrayerTime = praytime.CalculatePrayerTimes(&helsinki, time.Now())
	var torontoPrayerTime = praytime.CalculatePrayerTimes(&toronto, time.Now())
	var sandiegoPrayerTime = praytime.CalculatePrayerTimes(&sandiego, time.Now())
	var chicagoPrayerTime = praytime.CalculatePrayerTimes(&chicago, time.Now())

	printPrayerTimes("Helsinki", helsinkiPrayerTime)
	printPrayerTimes("Toronto", torontoPrayerTime)
	printPrayerTimes("San Diego", sandiegoPrayerTime)
	printPrayerTimes("Chicago", chicagoPrayerTime)

}

func printPrayerTimes(city string, prayertimes []string) {
	var today = time.Now()
	fmt.Println()
	fmt.Println("=======" + city + "=======")
	fmt.Println(today.Month(), today.Day(), today.Year())

	for x := 0; x < len(prayertimes); x++ {
		fmt.Println(praytime.PrayerTimeNames[x] + " - " + prayertimes[x])
	}

}
