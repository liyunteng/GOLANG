package main

import (
	"fmt"
	"log"
	"compress/gzip"
	"bytes"
	"os"
)

func gziptest() {
	buf := new(bytes.Buffer)
	gw := gzip.NewWriter(buf)

		var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		_, err := gw.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	err := gw.Close()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("test.gz", os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.Write(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	gr,err := gzip.NewReader(buf)
	if err != nil {
		log.Fatal(err)
	}
	defer gr.Close()

	out := make([]byte, 1024)
	for ; err == nil; _, err = gr.Read(out) {
		fmt.Printf("%s", out)
	}

}

func main() {
	gziptest()
}
