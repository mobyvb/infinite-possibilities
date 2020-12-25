package convertmany

import (
	"fmt"

	"github.com/mobyvb/temp-conversion/tempunit"
)

// Convert: convert many temperatures (byte form) from one unit to another
func Convert(from, to tempunit.TempUnit, values ...byte) {
	fmt.Println("convert many; starting values:")
	for i, v := range values {
		fmt.Printf("%d: %d\n", i, v)
	}

	celsiusValues := from.ConvertManyToCelsius(values)
	newValues := to.ConvertManyFromCelsius(celsiusValues...)

	for i, v := range values {
		fmt.Printf("%d: %d\t%d\n", i, v, newValues[i])
	}
}
