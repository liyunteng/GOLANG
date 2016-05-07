package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g probles.\n", math.Nextafter(1, 2))
}
