package convert3

import "fmt"

// Convert: types for each temp
func Convert(c, f, k, n float64) {
	fmt.Println("convert 3")
	cels := Celsius(c)
	fahr := Fahrenheit(f)
	kelv := Kelvin(k)
	newt := Newton(n)

	c2f := cels.ToFahrenheit()
	c2k := cels.ToKelvin()
	c2n := cels.ToNewton()

	f2c := fahr.ToCelsius()
	f2k := fahr.ToKelvin()
	f2n := fahr.ToNewton()

	k2f := kelv.ToFahrenheit()
	k2c := kelv.ToCelsius()
	k2n := kelv.ToNewton()

	n2c := newt.ToCelsius()
	n2f := newt.ToFahrenheit()
	n2k := newt.ToKelvin()

	fmt.Printf("%f deg C -> %f deg F, %f K, %f N\n", c, c2f, c2k, c2n)
	fmt.Printf("%f deg F -> %f deg C, %f K, %f N\n", f, f2c, f2k, f2n)
	fmt.Printf("%f K -> %f deg C, %f deg F, %f N\n", k, k2c, k2f, k2n)
	fmt.Printf("%f N -> %f deg C, %f deg F, %f K\n", n, n2c, n2f, n2k)
}

type Celsius float64
type Fahrenheit float64
type Kelvin float64
type Newton float64

func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit((9.0/5)*c + 32)
}

func (c Celsius) ToKelvin() Kelvin {
	return Kelvin(c + 273.15)
}

func (c Celsius) ToNewton() Newton {
	return Newton(c * 33.0 / 100)
}

func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((5.0 / 9) * (f - 32))
}

func (f Fahrenheit) ToKelvin() Kelvin {
	return f.ToCelsius().ToKelvin()
}

func (f Fahrenheit) ToNewton() Newton {
	return f.ToCelsius().ToNewton()
}

func (k Kelvin) ToCelsius() Celsius {
	return Celsius(k - 273.15)
}

func (k Kelvin) ToFahrenheit() Fahrenheit {
	return k.ToCelsius().ToFahrenheit()
}

func (k Kelvin) ToNewton() Newton {
	return k.ToCelsius().ToNewton()
}

func (n Newton) ToCelsius() Celsius {
	return Celsius(n * 100.0 / 33)
}

func (n Newton) ToFahrenheit() Fahrenheit {
	return n.ToCelsius().ToFahrenheit()
}

func (n Newton) ToKelvin() Kelvin {
	return n.ToCelsius().ToKelvin()
}
