package interfaces

import (
	"flag"
	"fmt"
)

// Celsius type
type Celsius float64

// Fahrenheit type
type Fahrenheit float64

// Constants
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF function
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC function
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type celsiusFlag struct {
	Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C":
		f.Celsius = Celsius(value)
		return nil
	case "F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines the Celsius flag with the specified name and value
// by default and the application instruction string and returns the address
// flag variable. The flag argument must contain a numeric value
// and the unit of measurement, for example "100 C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

// CelsiusConverter CLI
func CelsiusConverter() {
	var temp = CelsiusFlag("temp", 20.0, "temperature")
	flag.Parse()
	fmt.Println(*temp)
}
