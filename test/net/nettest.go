package main

import (
	"log"
	"net"
	"io"
)
func TestClient() {
	buf := make([]byte, 1024)
	for i := 0; i < 10; i++ {
		l, err := net.Dial("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}

		_, err = l.Write([]byte("hello, world!"))
		if err != nil {
			log.Fatal(err)
		}

		_, err = l.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		log.Printf("client recv: %s [%d]\n", buf, i)

	}
}
func TestListener() {
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			buf := make([]byte, 1024)
			_, err = c.Read(buf)
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
			log.Printf("server recv: %s\n", buf)
			_, err = c.Write(buf)
			if err != nil {
				log.Fatal(err)
			}
			c.Close()
		}(conn)

	}
}

func main() {
	go TestClient()
	TestListener()
}
