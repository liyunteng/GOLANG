package main

import (
	"bytes"
	"fmt"
	"os"
	"encoding/base64"
	"io"
	"sort"
)

func ExampleBuffer() {
	var b bytes.Buffer
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world!\n")

	b.WriteTo(os.Stdout)
}


func ExampleBuffer_reader() {
	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
}

func ExampleCompare() {
	var a, b []byte
	a = []byte("abcdefgh")
	b = []byte("ABCDEFG")
	b = bytes.ToLower(b)
	if bytes.Compare(a, b) < 0 {
		println("a<b")
	}
	if bytes.Compare(a, b) <= 0 {
		println("a<=b")
	}
	if bytes.Compare(a, b) > 0 {
		println("a>b")
	}
	if bytes.Compare(a, b) >= 0 {
		println("a>=b")
	}

	if bytes.Equal(a, b) {
		println("a==b")
	}
	if !bytes.Equal(a, b) {
		println("a!=b")
	}
}

func ExampleCompare_search() {
	var needle []byte = []byte("abc")
	haystack := [][]byte{[]byte("abcdefg"), []byte("hijklmn"), []byte("opq")}

	i := sort.Search(len(haystack), func(i int) bool {
		return bytes.Compare(haystack[i], needle) >= 0
	})
	if i < len(haystack) && bytes.Equal(haystack[i], needle) {
		println("found it")
	} else {
		println("not found")
	}
}

func ExampleTrimSuffix() {
	var b = []byte("Hello, goodbye, etc!")
	b = bytes.TrimSuffix(b, []byte("goodbye, etc!"))
	b = bytes.TrimSuffix(b, []byte("gopher"))
	b = append(b, bytes.TrimSuffix([]byte("world!x!x!"), []byte("x!")) ...)
	os.Stdout.Write(b)
}

func ExampleTrimPrefix() {
	var b = []byte("Goodbye,, world!")
	b = bytes.TrimPrefix(b,[]byte("Goodbye,"))
	b = bytes.TrimPrefix(b, []byte("See ya,"))
	println("Hello" + string(b))
}

func main() {
	ExampleBuffer()
	ExampleBuffer_reader()
	ExampleCompare()
	ExampleCompare_search()
	ExampleTrimSuffix()
	ExampleTrimPrefix()
}
