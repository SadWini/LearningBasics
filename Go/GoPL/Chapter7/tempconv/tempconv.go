package tempconv

import (
	"flag"
	"fmt"
)

type celsiusFlag struct{ Celsius }

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
	case "K", "°K":
		f.Celsius = KToC(Kelvin(value))
	}
	return fmt.Errorf("invalid temperature %q", s)
}

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsolutZeroC Celsius = -273.15
	FreezingC            = 0
	BoilingC             = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g*C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g*F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%g*K", k)
}

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}
