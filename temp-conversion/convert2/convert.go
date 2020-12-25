package convert2

import "fmt"

// Convert: basic functions
func Convert(c, f, k, n float64) {
	fmt.Println("convert 2")
	c2f := celsiusToFahrenheit(c)
	c2k := celsiusToKelvin(c)
	c2n := celsiusToNewton(c)

	f2c := fahrenheitToCelsius(f)
	f2k := fahrenheitToKelvin(f)
	f2n := fahrenheitToNewton(f)

	k2f := kelvinToFahrenheit(k)
	k2c := kelvinToCelsius(k)
	k2n := kelvinToNewton(k)

	n2c := newtonToCelsius(n)
	n2f := newtonToFahrenheit(n)
	n2k := newtonToKelvin(n)

	fmt.Printf("%f deg C -> %f deg F, %f K, %f N\n", c, c2f, c2k, c2n)
	fmt.Printf("%f deg F -> %f deg C, %f K, %f N\n", f, f2c, f2k, f2n)
	fmt.Printf("%f K -> %f deg C, %f deg F, %f N\n", k, k2c, k2f, k2n)
	fmt.Printf("%f N -> %f deg C, %f deg F, %f K\n", n, n2c, n2f, n2k)
}

func fahrenheitToKelvin(f float64) float64 {
	return celsiusToKelvin(fahrenheitToCelsius(f))
}

func fahrenheitToCelsius(f float64) float64 {
	return (5.0 / 9) * (f - 32)
}

func fahrenheitToNewton(f float64) float64 {
	return celsiusToNewton(fahrenheitToCelsius(f))
}

func celsiusToFahrenheit(c float64) float64 {
	return (9.0/5)*c + 32
}

func celsiusToKelvin(c float64) float64 {
	return c + 273.15
}

func celsiusToNewton(c float64) float64 {
	return c * 33.0 / 100
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func kelvinToNewton(k float64) float64 {
	return celsiusToNewton(kelvinToCelsius(k))
}

func newtonToCelsius(n float64) float64 {
	return n * 100.0 / 33
}

func newtonToFahrenheit(n float64) float64 {
	return celsiusToFahrenheit(newtonToCelsius(n))
}

func newtonToKelvin(n float64) float64 {
	return celsiusToKelvin(newtonToCelsius(n))
}
