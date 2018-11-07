package main

import (
	"flag"
	"fmt"
	"time"
)


func main() {
	boolv := flag.Bool("test_bool", false, "bool value")
	intv :=flag.Int("test_int", 10, "int value")
	int64v := flag.Int64("test_int64", 10, "int64 value")
	uintv := flag.Uint("test_uint", 10, "uint value")
	uint64v := flag.Uint64("test_uint64", 10, "uint64 value")
	stringv := flag.String("test_string", "", "string value")
	float64v := flag.Float64("test_float64", 10.0, "float64 value")
	durationv := flag.Duration("test_duration", 1e10, "time.Duration value")

	flag.Parse()

	fmt.Println(*boolv)
	fmt.Println(*intv)
	fmt.Println(*int64v)
	fmt.Println(*uintv)
	fmt.Println(*uint64v)
	fmt.Println(*stringv)
	fmt.Println(*float64v)
	fmt.Println(*durationv)


	fmt.Println()
	for i, param := range flag.Args() {
		fmt.Printf("%d : %v\n", i, param)
	}
	flag.Usage()

	fmt.Printf("%T: %v", durationv, *durationv)
	time.Sleep(*durationv)

}
