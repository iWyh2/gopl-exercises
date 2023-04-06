package ch2

type Kelvin float64
type Celsius float64

const (
	AbsoluteZeroK Kelvin = 0
)

func KToC(k Kelvin) Celsius {
	return Celsius(k + 273.15)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c - 273.15)
}
