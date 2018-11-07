package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct {}

func (h Hello) ServeHTTP (
	w http.ResponseWriter,
	r *http.Request) {
		fmt.Fprint(w, "hello!")
	}
type String string
func (s String) String() string {
	return string(s)
}
func (s String) ServeHTTP (
	w http.ResponseWriter,
	r *http.Request) {
		fmt.Fprintf(w, s.String())
	}

type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (s Struct) ServeHTTP (
	w http.ResponseWriter,
	r *http.Request) {
		fmt.Fprintf(w, s.Greeting + s.Punct + s.Who)
	}
func main() {
	http.Handle("/string", String("I'm a good boy!"))
	http.Handle("/struct", &Struct{"Hello", ":", "lyt"})
	http.ListenAndServe(":8080", nil)
	var h Hello
	err := http.ListenAndServe("localhost:4000", h)
	if err != nil {
		log.Fatal(err)
	}

}
