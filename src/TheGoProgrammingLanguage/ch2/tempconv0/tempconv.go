package tempconv

import (
	"fmt"
)

type Celsius float64
type Fahreheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahreheit { return Fahreheit(c*9/5 + 32) }
func FToC(f Fahreheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string   { return fmt.Sprintf("%g°C", c) }
func (f Fahreheit) String() string { return fmt.Sprintf("%g°F", f) }
