package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ExampleWriter() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	w.Flush()
}

func ExampleScanner_lines() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading stanard input:", err)
	}
}

func ExampleScanner_words() {
	const input = "Now is the winter of our discontent,\nMade glorious summer bye this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Printf("%d\n", count)
}

func ExampleScanner_custom() {
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	split := func(data[]byte, atEOF bool)(advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}

	scanner.Split(split)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
}


func main() {
	ExampleWriter()
	ExampleScanner_lines()
	ExampleScanner_words()
	ExampleScanner_custom()

}
