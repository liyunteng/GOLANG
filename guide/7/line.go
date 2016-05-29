package main

import (
	"fmt"
)


type Vector []int64

func main() {
	u := make(Vector, 80960000)

	for i, _ := range u {
		u[i] = int64(i)+1
	}

	var result int64
	for _, v := range u {
		result += v
	}

	fmt.Println("result:", result)
}
