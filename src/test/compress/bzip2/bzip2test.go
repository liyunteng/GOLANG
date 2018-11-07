package main

import (
	"bytes"
	"io/ioutil"
	"compress/bzip2"
	"log"
	"fmt"
	// "io"
	// "os"
)

func TestBitReader(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("content: %v\n", string(

	br := bzip2.NewReader(bytes.NewReader(content))
	// if _, err = io.Copy(os.Stdout, br); err != nil {
	//	log.Fatal(err)
	// }


	buf := make([]byte, 512)
	for ;  err == nil ; _,err = br.Read(buf) {
		fmt.Printf("%s", buf)
	}

}

func main() {
	TestBitReader("test.bz2")
}
