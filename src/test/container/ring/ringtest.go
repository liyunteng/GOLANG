package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(10)
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	r = r.Move(3)
	r = r.Unlink(r.Len()-1)
	for i := 0; i < r.Len(); i++{
		fmt.Printf("%.2d:%d\n", i+1, r.Value)
		r = r.Next()
	}

	s := 0
	r.Do(func (p interface{}) {
			if p != nil {
				s += p.(int)
			}
	})

	fmt.Printf("sum: %v\n", s)
}
