// parse test.xml
package main

import (
	"fmt"
	"io/ioutil"
	"encoding/xml"
)

const (
	xmlFile string = "test.xml"
)

type Root struct {
	XMLName xml.Name `xml:"root"`
	DevId string `xml:"dev_id"`
	UserId string `xml:"user_id"`
	VspUrl string `xml:"vsp_url"`
	Code int `xml:"code"`
	Msg string `xml:"msg"`
	User User `xml:"user"`
	Contacts Contacts `xml:"contacts"`
	Msms Msms `xml:"msms"`
}

type User struct {
	UserId string `xml:"user_id"`
	UserAccount string `xml:"user_account"`
	UserName string `xml:"user_name"`
	BigVisitNum int `xml:"big_visit_num"`
	VspId string `xml:"vsp_id"`
	VspName string `xml:"vsp_name"`
}

type Contacts struct {
	Contact []Contact `xml:"contact"`
}

type Contact struct {
	UserId string `xml:"user_id"`
	UserAccount string `xml:"user_account"`
	UserName string `xml:"user_name"`
}

type Msms struct {
	Msm []Msm `xml:"msm"`
}

type Msm struct {
	MsmIp string `xml:"msm_ip"`
	MsmPort int `xml:"msm_port"`
}

// parse file
func parseXml(file string) error {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("readfile %q failed: %v\n", file, err)
		return err
	}

	var root Root

	err = xml.Unmarshal(content, &root)
	if err != nil {
		fmt.Printf("parse %q failed: %v\n", file, err)
		return err
	}

	fmt.Println(root)
/*
	msm := Msm{"12.12.12.12", 1000}

	root.Msms.Msm = append(root.Msms.Msm[:], msm)

	output, err := xml.MarshalIndent(root, "", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return err
	}
*/
	return nil
}

func main() {
	parseXml(xmlFile)
}
