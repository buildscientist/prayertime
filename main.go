package main

import "math"
import "fmt"
import "github.com/buildscientist/prayertime/trig"


func main() {
	fmt.Println("Testing 123")
	fmt.Println(math.Sin(trig.DegreesToRadians(90)))	
}

//Prayer Time Calculation functions
func sunPosition(julianDate float64) (position []float64) {

	var daysFromEpoch = julianDate - 2451545.0

	return []float64{daysFromEpoch, 2.0, 3.0}
}
