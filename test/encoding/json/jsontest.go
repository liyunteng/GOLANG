package main

import (
	"fmt"
	"encoding/json"
	"os"
	"strings"
	"io"
	"log"
	"bytes"
)

func TestMarshal() {
	type ColorGroup struct {
		ID int
		Name string
		Colors []string
	}

	group := ColorGroup{
		ID:	1,
		Name:	"Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")
	os.Stdout.Write(out.Bytes())
}

func TestUnmarshal() {
	var jsonBolb = []byte(`[
{"Name": "Platypus", "Order": "Monotremata"},
{"Name": "Quoll", "Order": "Dasyuromorphia"}
]`)

	type Animal struct {
		Name string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBolb, &animals)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Printf("%#v", animals)
}

func TestDecoder() {
	const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
		{"Name": "Sam", "Text": "Who's there?"}
		{"Name": "Ed", "Text": "Go fmt."}
		{"Name": "Sam", "Text": "Go fmt who?"}
		{"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}

func TestDecoder_Token() {
	const jsonStream = `{"Message": "Hello", "Array":[1, 2, 3], "Null": null, "Number":1.234}`

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		t, err := dec.Token()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%T: %v", t, t)
		if dec.More() {
			fmt.Printf(" (more)")
		}
		fmt.Printf("\n")
	}
}
func main() {
	TestMarshal()
	TestUnmarshal()
	TestDecoder()
	TestDecoder_Token()
}
