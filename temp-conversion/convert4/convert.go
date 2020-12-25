package convert4

import (
	"fmt"

	"github.com/mobyvb/temp-conversion/tempunit"
)

// Convert: agnostic type for temperature
func Convert(c, f, k, n float64) {
	fmt.Println("convert 4")
	cels := tempFromCelsius(c)
	fahr := tempFromFahrenheit(f)
	kelv := tempFromKelvin(k)
	newt := tempFromNewton(n)

	c2f := cels.Fahrenheit()
	c2k := cels.Kelvin()
	c2n := cels.Newton()

	f2c := fahr.Celsius()
	f2k := fahr.Kelvin()
	f2n := fahr.Newton()

	k2f := kelv.Fahrenheit()
	k2c := kelv.Celsius()
	k2n := kelv.Newton()

	n2f := newt.Fahrenheit()
	n2c := newt.Celsius()
	n2k := newt.Kelvin()

	fmt.Printf("%f deg C -> %f deg F, %f K, %f N\n", c, c2f, c2k, c2n)
	fmt.Printf("%f deg F -> %f deg C, %f K, %f N\n", f, f2c, f2k, f2n)
	fmt.Printf("%f K -> %f deg C, %f deg F, %f N\n", k, k2c, k2f, k2n)
	fmt.Printf("%f N -> %f deg C, %f deg F, %f K\n", n, n2c, n2f, n2k)
}

// Temp is a general temperature value (value stored is in celsius)
type Temp float64

func tempFromCelsius(c float64) Temp {
	return Temp(c)
}

func tempFromFahrenheit(f float64) Temp {
	c := tempunit.F.ConvertToCelsius(f)
	return Temp(c)
}

func tempFromKelvin(k float64) Temp {
	c := tempunit.K.ConvertToCelsius(k)
	return Temp(c)
}

func tempFromNewton(n float64) Temp {
	c := tempunit.N.ConvertToCelsius(n)
	return Temp(c)
}

func (t Temp) Celsius() float64 {
	return float64(t)
}

func (t Temp) Fahrenheit() float64 {
	return tempunit.F.ConvertFromCelsius(float64(t))
}

func (t Temp) Kelvin() float64 {
	return tempunit.K.ConvertFromCelsius(float64(t))
}

func (t Temp) Newton() float64 {
	return tempunit.N.ConvertFromCelsius(float64(t))
}
