package main

import (
	"fmt"
	"github.com/buildscientist/prayertime/praytime"
	"time"
)

func main() {
	var chicago = praytime.New(41.879626, -87.648217, -5.0)
	var sandiego = praytime.New(32.7157, -117.1611, -7.0)

	chicago.TimeFormat = praytime.TIME_24
	sandiego.TimeFormat = praytime.TIME_24

	var sandiegoPrayerTime = praytime.CalculatePrayerTimes(&sandiego, time.Now())
	var chicagoPrayerTime = praytime.CalculatePrayerTimes(&chicago, time.Now())

	printPrayerTimes("San Diego", sandiegoPrayerTime)
	printPrayerTimes("Chicago", chicagoPrayerTime)

}

func printPrayerTimes(city string, prayertimes []string) {
	var today = time.Now()
	fmt.Println()
	fmt.Println(today.Month(), today.Day(), today.Year())
	fmt.Println("============" + city + "============")

	for x := 0; x < len(prayertimes); x++ {
		fmt.Println(praytime.PrayerTimeNames[x] + " - " + prayertimes[x])
	}

}
