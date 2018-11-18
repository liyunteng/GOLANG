package tempconv

import "fmt"
import "testing"

func TestOne(t *testing.T) {
	//!+arith
	fmt.Printf("%g\n", BoilingC - FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF - CToF(FreezingC)) // "180" °F
	//!-arith
}

func TestTwo(t *testing.T) {
	//!+printf
	c := FToC(212.0)
	fmt.Println(c.String())		// "100°C"
	fmt.Printf("%v\n", c)		// "100°C"
	fmt.Printf("%s\n", c)		// "100°C"
	fmt.Println(c)				// "100°C"
	fmt.Printf("%g\n", c)		// "100"
	fmt.Println(float64(c))		// "100"
	//!-printf
}
