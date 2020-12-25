package tempunit

// TempUnit defines the scale and offset necessary for converting a temperature in celsius to a particular unit.
type TempUnit struct {
	scale  float64
	offset float64
}

var (
	F = TempUnit{
		scale:  9 / 5.0,
		offset: 32,
	}
	K = TempUnit{
		scale:  1,
		offset: 273.15,
	}
	N = TempUnit{
		scale:  33.0 / 100,
		offset: 0,
	}
	C = TempUnit{
		scale:  1,
		offset: 0,
	}
)

func (from TempUnit) ConvertToCelsius(value float64) (celsius float64) {
	return (value - from.offset) / from.scale
}

func (to TempUnit) ConvertFromCelsius(celsius float64) (newValue float64) {
	return to.scale*celsius + to.offset
}

func (from TempUnit) ConvertManyToCelsius(values ...byte) (celsiusValues []byte) {
	return []byte{}
}

func (to TempUnit) ConvertManyFromCelsius(values ...byte) (newValues []byte) {
	return []byte{}
}
