// Package tempconv performs Celsius, Kelvin and Fahrenheit conversions
package tempconv

import "fmt"

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
