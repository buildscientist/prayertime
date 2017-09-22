/*
Copyright 2017
Author: Youssuf ElKalay
All rights reserved. Use of this source code is governed by the Apache 2.0 license
that can be found in the LICENSE file or at http://www.apache.org/licenses/LICENSE-2.0
*/

package praytime

//Prayer Times
const FAJR string = "fajr"
const SUNRISE string = "sunrise"
const DHUHR string = "dhuhr"
const ASR string = "asr"
const SUNSET = "sunset"
const MAGHRIB string = "maghrib"
const ISHA string = "isha"

//Enumeration of prayer names with their corresponding element IDs in praytime.CalculatePrayerTimes
const (
	FAJR_TIME    = 0
	SUNRISE_TIME = 1
	DHUHR_TIME   = 2
	ASR_TIME     = 3
	SUNSET_TIME  = 4
	MAGHRIB_TIME = 5
	ISHA_TIME    = 6
)

//Location based prayer calculation methods
const JAFARI int = 0
const KARACHI int = 1
const ISNA int = 2
const MWL int = 3
const MAKKAH int = 4
const EGYPT int = 5
const TEHRAN int = 6
const CUSTOM int = 7

//Asr Juristic Methods
const SHAFII int = 0
const HANAFI int = 1
const DHUHR_MINUTES int = 0

//Adjustments for higher altitudes
const NONE int = 0
const MIDNIGHT int = 1
const ONE_SEVENTH int = 2
const ANGLE_BASED int = 3

//Time Formats
const TIME_24 int = 0
const TIME_12 int = 1
const TIME_12_NO_SUFFIX int = 2

const NUMBER_OF_ITERATIONS int = 1
const INVALID_TIME string = "-----"
