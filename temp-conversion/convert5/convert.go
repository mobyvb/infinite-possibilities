package convert5

import (
	"fmt"

	"github.com/mobyvb/temp-conversion/tempunit"
)

// Convert: single function for all conversions, specify from unit and to unit
func Convert(c, f, k, n float64) {
	fmt.Println("convert 5")
	c2f := convertTemp(c, tempunit.C, tempunit.F)
	c2k := convertTemp(c, tempunit.C, tempunit.K)
	c2n := convertTemp(c, tempunit.C, tempunit.N)

	f2c := convertTemp(f, tempunit.F, tempunit.C)
	f2k := convertTemp(f, tempunit.F, tempunit.K)
	f2n := convertTemp(f, tempunit.F, tempunit.N)

	k2f := convertTemp(k, tempunit.K, tempunit.F)
	k2c := convertTemp(k, tempunit.K, tempunit.C)
	k2n := convertTemp(k, tempunit.K, tempunit.N)

	n2f := convertTemp(n, tempunit.N, tempunit.F)
	n2c := convertTemp(n, tempunit.N, tempunit.C)
	n2k := convertTemp(n, tempunit.N, tempunit.K)

	fmt.Printf("%f deg C -> %f deg F, %f K, %f N\n", c, c2f, c2k, c2n)
	fmt.Printf("%f deg F -> %f deg C, %f K, %f N\n", f, f2c, f2k, f2n)
	fmt.Printf("%f K -> %f deg C, %f deg F, %f N\n", k, k2c, k2f, k2n)
	fmt.Printf("%f N -> %f deg C, %f deg F, %f K\n", n, n2c, n2f, n2k)
}

func convertTemp(value float64, from, to tempunit.TempUnit) float64 {
	c := from.ConvertToCelsius(value)
	return to.ConvertFromCelsius(c)
}
