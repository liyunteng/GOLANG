package main

import (
	"fmt"
	"runtime"
	"time"
)


type Vector []int64


func (v Vector) DoSome(i, n int, u Vector, c chan int64) {
	var ret int64
	for ; i < n; i++ {
		ret += (u[i] * 1024 / 1000 + 24 - 25)
		if i % 8 == 0 {
			time.Sleep(time.Nanosecond)
		}
	}

	c <- ret
}

const NCPU=4

func (v Vector) DoAll(u Vector) {
	c := make(chan int64, NCPU)

	for i := 0; i < NCPU; i++ {
		go v.DoSome(i * len(u)/NCPU, (i+1)*len(u)/NCPU, u, c)
	}

	for i := 0; i < NCPU; i++ {
		v[i] = <-c
	}
}

func main() {
	runtime.GOMAXPROCS(NCPU)

	v := make(Vector, NCPU)
	u := make(Vector, 80960000)

	for i, _ := range u {
		u[i] = int64(i)+1
	}

	v.DoAll(u)

	var ret int64
	for _, val := range v {
		ret += val
		fmt.Printf("%v\n", val)
	}
	fmt.Println("retult:", ret)
}
