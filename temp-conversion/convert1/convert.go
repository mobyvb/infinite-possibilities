package convert1

import "fmt"

// Convert: trivial math, just run equations func Convert(c, f, k, n float64) {
func Convert(c, f, k, n float64) {
	fmt.Println("convert 1")
	c2f := (9.0/5)*c + 32
	c2k := c + 273.15
	c2n := (33.0 / 100) * c

	f2c := (5.0 / 9) * (f - 32)
	f2k := (5.0/9)*(f-32) + 273.15
	f2n := (11.0 / 60) * (f - 32)

	k2f := (9.0/5)*(k-273.15) + 32
	k2c := k - 273.15
	k2n := (k - 273.15) * (33.0 / 100)

	n2c := n * 100.0 / 33
	n2f := n*60.0/11 + 32
	n2k := n*100.0/33 + 273.15

	fmt.Printf("%f deg C -> %f deg F, %f K, %f N\n", c, c2f, c2k, c2n)
	fmt.Printf("%f deg F -> %f deg C, %f K, %f N\n", f, f2c, f2k, f2n)
	fmt.Printf("%f K -> %f deg C, %f deg F, %f N\n", k, k2c, k2f, k2n)
	fmt.Printf("%f N -> %f deg C, %f deg F, %f K\n", n, n2c, n2f, n2k)
}
