package convert6

import (
	"fmt"

	"github.com/mobyvb/temp-conversion/tempunit"
)

// Convert: single function for all conversions, returns all units
func Convert(c, f, k, n float64) {
	fmt.Println("convert  6")
	c2f, _, c2k, c2n := convertTemp2(c, tempunit.C)

	_, f2c, f2k, f2n := convertTemp2(f, tempunit.F)

	k2f, k2c, _, k2n := convertTemp2(k, tempunit.K)

	n2f, n2c, n2k, _ := convertTemp2(n, tempunit.N)

	fmt.Printf("%f deg C -> %f deg F, %f K, %f N\n", c, c2f, c2k, c2n)
	fmt.Printf("%f deg F -> %f deg C, %f K, %f N\n", f, f2c, f2k, f2n)
	fmt.Printf("%f K -> %f deg C, %f deg F, %f N\n", k, k2c, k2f, k2n)
	fmt.Printf("%f N -> %f deg C, %f deg F, %f K\n", n, n2c, n2f, n2k)
}

func convertTemp2(value float64, from tempunit.TempUnit) (f, c, k, n float64) {
	c = from.ConvertToCelsius(value)

	f = tempunit.F.ConvertFromCelsius(c)
	k = tempunit.K.ConvertFromCelsius(c)
	n = tempunit.N.ConvertFromCelsius(c)

	return f, c, k, n
}
