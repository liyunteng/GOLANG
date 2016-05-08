package main

import (
	"archive/zip"
	"bytes"
	// "compress/flate"
	"fmt"
	"io"
	"log"
	"os"
)

func ExampleWriter() {
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "This archive contains some text files."},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}

	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}

	}

	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	v, err := os.OpenFile("test.zip", os.O_CREATE | os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
	_, err = v.Write(buf.Bytes())
		if err != nil {
			log.Fatal(err)
		}
		v.Close()
}

func ExampleReader() {
	r, err := zip.OpenReader("test.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, f.FileInfo().Size())
		if err  != nil {
			log.Fatal(err)
		}

		rc.Close()
		fmt.Println()
	}
}

func main() {
	ExampleWriter()
	ExampleReader()
}
