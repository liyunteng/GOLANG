package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func Example() {
	buf := new(bytes.Buffer)

	tw := tar.NewWriter(buf)

	var files = [] struct {
		Name, Body string
	} {
		{"readme.txt", "This archive contails some text files."},
		{"lyt.txt", "lyt names:\nLi\nYunteng"},
		{"todo.txt", "Get animal handing license."},
	}

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}

		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}
	}

	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

	f, err := os.OpenFile("test.tar", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = f.Write(buf.Bytes())
	if err != nil {
		log.Fatalln(err)
	}
	f.Close()

	r := bytes.NewReader(buf.Bytes())
	tr := tar.NewReader(r)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Contents of %s:\n", hdr.Name)
		if _, err := io.Copy(os.Stdout, tr); err != nil {
			log.Fatalln(err)
		}
		fmt.Println()
	}

}

func main() {
	Example()
}
