package convert7

import (
	"fmt"

	"github.com/mobyvb/temp-conversion/tempunit"
)

// Convert: agnostic temperature class that does conversions in constructor
func Convert(c, f, k, n float64) {
	fmt.Println("convert 7")
	t1 := newTemp(c, tempunit.C)
	c2f := t1.Fahrenheit
	c2k := t1.Kelvin
	c2n := t1.Newton

	t2 := newTemp(f, tempunit.F)
	f2c := t2.Celsius
	f2k := t2.Kelvin
	f2n := t2.Newton

	t3 := newTemp(k, tempunit.K)
	k2f := t3.Fahrenheit
	k2c := t3.Celsius
	k2n := t3.Newton

	t4 := newTemp(n, tempunit.N)
	n2f := t4.Fahrenheit
	n2c := t4.Celsius
	n2k := t4.Kelvin

	fmt.Printf("%f deg C -> %f deg F, %f K, %f N\n", c, c2f, c2k, c2n)
	fmt.Printf("%f deg F -> %f deg C, %f K, %f N\n", f, f2c, f2k, f2n)
	fmt.Printf("%f K -> %f deg C, %f deg F, %f N\n", k, k2c, k2f, k2n)
	fmt.Printf("%f N -> %f deg C, %f deg F, %f K\n", n, n2c, n2f, n2k)
}

type Temp struct {
	Fahrenheit float64
	Celsius    float64
	Kelvin     float64
	Newton     float64
}

func newTemp(value float64, unit tempunit.TempUnit) Temp {
	var f, c, k, n float64

	c = unit.ConvertToCelsius(value)

	f = tempunit.F.ConvertFromCelsius(c)
	k = tempunit.K.ConvertFromCelsius(c)
	n = tempunit.N.ConvertFromCelsius(c)

	return Temp{
		Fahrenheit: f,
		Celsius:    c,
		Kelvin:     k,
		Newton:     n,
	}
}
