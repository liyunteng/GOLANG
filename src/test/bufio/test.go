package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		v,err := strconv.ParseInt(string(scanner.Text()), 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("value: %+v %T\n", v, v)
	}
	// split := func(data[]byte, atEOF bool)(advance int, token []byte, err error) {
	//	advance, token, err = bufio.ScanWords(data, atEOF)
	//	fmt.Printf("token: %+v\n", string(token))
	//	if err == nil && token != nil {
	//		_, err = strconv.ParseInt(string(token), 10, 32)
	//	}
	//	return
	// }

	// scanner.Split(split)
	// for scanner.Scan() {
	//	fmt.Printf("value: %+v %T\n", scanner.Text(), scanner.Text())
	// }

	if err := scanner.Err(); err != nil {
		fmt.Printf("scan err: %v\n", err)
	}



}
