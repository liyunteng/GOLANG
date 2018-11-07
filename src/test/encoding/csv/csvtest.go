package main

import (
	"fmt"
	"log"
	"encoding/csv"
	"strings"
	"io"
	"os"
)

func TestReader() {
	in := `first_name, last_name, username
"Rob","Pike", rob
Ken, Thompson, ken
"Robert","Griesemer","gri"`

	r := csv.NewReader(strings.NewReader(in))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}

func TestReader_options() {
		in := `first_name;last_name;username
"Rob";"Pike";rob
# lines beginning with a # character are ignored
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ';'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
}

func TestWriter() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)

	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func TestWriteAll() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv: ", err)
	}
}
func main() {
	TestReader()
	TestReader_options()
	TestWriter()
	TestWriteAll()
}
