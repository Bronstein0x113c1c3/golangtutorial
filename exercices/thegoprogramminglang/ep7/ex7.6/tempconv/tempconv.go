package tempconv

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

/*

package flag
// Value is the interface to the value stored in a flag.
type Value interface {
String() string
Set(string) error
}
*/
