package  main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"log"
	"os"
)

func zlibtest() {
	buf := new(bytes.Buffer)
	zw := zlib.NewWriter(buf)

			var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		_, err := zw.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	err := zw.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	zr, err := zlib.NewReader(buf)
	if err != nil {
		log.Fatal(err)
	}
	defer zr.Close()

	out := make([]byte, 1024)
	for ; err == nil; _, err = zr.Read(out) {
		fmt.Printf("%s", out)
	}
}

func main() {
	zlibtest()
}
