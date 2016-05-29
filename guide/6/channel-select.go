package main

import "fmt"

func main() {
	ch := make(chan string, 1)

	for {
		select {
		case ch <- "nihao":
		case ch <- "abcdefghik##########33":
		}

		i := <-ch
		fmt.Println("received: ", i)
	}
}
