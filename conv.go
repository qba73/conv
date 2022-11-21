package conv

import "fmt"

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	Boiling       Celsius = 100
)

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type Fahrenheit float64

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

// CToF converts a Celsius temp to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a Fahrenheit temp to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
