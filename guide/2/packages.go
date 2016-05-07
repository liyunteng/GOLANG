package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n := 10;
	rand.Seed(int64(time.Now().Nanosecond()))
	fmt.Println("My favorite number is", rand.Intn(n))
	fmt.Println("My favorite number is", rand.Intn(n))
	fmt.Println("My favorite number is", rand.Intn(n))
	fmt.Println("My favorite number is", rand.Intn(n))
}
