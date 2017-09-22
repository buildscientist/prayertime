# PrayTime
PrayTime is a Go based library that calculates Muslim prayer times occurring 5 times a day. The library was ported from Hamid Zarrabi-Zadeh and Hussain Ali Khan's [Java](http://praytimes.org/code/git/?a=viewblob&p=PrayTimes&h=093f77d6cc83b53fb12e9900803d5fa75dacd110&hb=HEAD&f=v1/java/PrayTime.java) implementation. 

Full details as to the overall prayer calculation algorithm are available on the [Praytimes.org](http://praytimes.org/calculation) site.


---
## Prerequisites 
1. Go (v1.7 of above) installed on target workstation.
1. [GOPATH](https://github.com/golang/go/wiki/Setting-GOPATH) environment variable set correctly.
---
## Installation 

```go
go get github.com/buildscientist/prayertime 
```

---

## API
The PrayTime Go API is fairly simple to use - the user does not need to understand the underlying calculations to use the library. 

The PrayerLocale struct is defined as follows: 

```go
type PrayerLocale struct {
	latitude, longitude, timezone float64
	PrayerCalcMethod,AsrJuristic,AdjustHighLats,TimeFormat int
}
```

It requires 3 parameters: 
- Latitude - integer value from 90 to -90
- Longitude - integer value from -180 to 180
- Timezone - integer value denoting offset from UTC 

The other parameters are not required and are preset at instantiation including: 

- PrayerCalcMethod - set to praytime.ISNA
- AsrJuristic - set to praytime.SHAFII by default
- AdjustHighLats - set to praytime.NONE by default. 
- TimeFormat - set to praytime.TIME_12 by default. Displays prayer times in the format of hh:mm suffix 

---
## Example 
```go
package main 

import (
     "fmt"
     "time"
     "github.com/buildscientist/prayertime/praytime"
) 

	name,offset := time.Now().Zone()
	name = name 
	timezone = float64(offset/3600)

	var chicago = praytime.New(41.879626, -87.648217,timezone)
	var chicagoPrayerTime = praytime.CalculatePrayerTimes(&chicago, time.Now())

func printPrayerTimes(city string, prayertimes []string) {
 	var today = time.Now()
	fmt.Println()
	fmt.Println("=======" + city + "=======")
	fmt.Println(today.Month(), today.Day(), today.Year())

	for x := 0; x < len(prayertimes); x++ {
		fmt.Println(praytime.PrayerTimeNames[x] + " - " + prayertimes[x])
	}

}
```

This should output something similar to the following: 

```json
=======Chicago=======
September 20 2017
fajr - 05:18 AM
sunrise - 06:35 AM
dhuhr - 12:44 PM
asr - 04:14 PM
sunset - 06:55 PM
maghrib - 06:55 PM
isha - 08:12 PM

```
