// xmlReader to parse layout.xml
package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
)

const (
	xmlFile string = "layout.xml"
)

type Root struct {
	XMLName xml.Name `xml:"root"`
	Layouts []Layout `xml:"layout"`
}

type Layout struct {
	Screens int `xml:"screens"`
	Current string `xml:"current"`
	Modes []Mode `xml:"mode"`
}

type Mode struct {
	Name string `xml:"name"`
	Screens []Screen `xml:"screen"`
}

type Screen struct {
	StartX int `xml:"startX"`
	StartY int `xml:"startY"`
	Width int `xml:"width"`
	Height int `xml:"height"`
	ScreenId int `xml:"screenId"`
}

func parseXml (file string) error {
	content,err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("open %q failed: %v\n", file, err)
		return err
	}

	var root Root
	err = xml.Unmarshal(content, &root)
	if err != nil {
		fmt.Printf("parse %q failed: %v\n", file, err)
		return err
	}

	fmt.Println(root)
	return nil
}

func main() {
	parseXml(xmlFile)
}
