package praytime

type PrayTime struct {
	latitude,longitude,timezone float64
}

func New(latitude,longitude,timezone float64) PrayTime {
	return PrayTime{latitude,longitude,timezone}
}

func Foo() string {
	return "Does nothing\n"
}