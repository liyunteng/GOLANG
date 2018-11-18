package main

import (
	"fmt"
)

func main() {
	in := make(chan int)
	out := make(chan int)
	go counter(in)
	go squarer(out, in)
	printer(out)
}

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out<- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
