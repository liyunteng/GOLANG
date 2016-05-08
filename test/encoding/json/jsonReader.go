package main

import (
	"fmt"
	"log"
	"encoding/json"
	"bytes"
)

func TestService() {
	var jsons []interface{}
	type Test struct {
		Action string `json:"action"`
	}
	test := Test{"Test"}
	jsons = append(jsons, test)

	type SetFBTransparent struct {
		Action string `json:"action"`
		Value  uint64 `json:"value"`
	}
	fbt := SetFBTransparent {"setFBTransparent", 1056816}
	jsons = append(jsons, fbt)

	type LrAddr [2]interface{}
	type Start struct {
		Action string `json:"action"`
		LocalCamUrl string `json:"localCamUrl"`
		LrAddr LrAddr `json:"lrAddr"`
		MyName string `json:"myName"`
		MySessionNum uint `json:"mySeesionNum"`
		ParticipantNameList interface{} `json:"participantNameList"`
		Type string `json:"type"`
		UpstreamingBandwidth uint `json:"upstreamingBandwidth"`
	}
	start := Start {"start", "rtsp://192.0.0.64:554/1", LrAddr{"172.16.1.219", 60615}, "fot test", 1, nil, "cmd", 20000}
	jsons = append(jsons, start)

	type Picture struct {
		StartX uint `json:"startX"`
		StartY uint `json:"startY"`
		Height uint `json:"height"`
		Width uint `json:"width"`
		Volume uint `json:"volume"`
		StreamId uint `json:"streamId"`
	}
	type AvLayout struct {
		Action string `json:"action"`
		Pictures [4]Picture `json:"pictures"`
		ScreenId uint `json:"screenId"`
	}
	layout := AvLayout {"avSetLayout", [4]Picture{{0, 0, 1080, 1920, 0, 0}, {1280, 720, 360, 640, 1, 65535}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}}, 0}
	jsons = append(jsons, layout)

	for _, v := range(jsons) {
		b, err := json.Marshal(v)
		if err != nil {
			log.Fatal(err)
		}
		var buf bytes.Buffer
		json.Indent(&buf, b, "", "\t")
		fmt.Printf("%v\n", string(buf.Bytes()))

		var data map[string]interface{}
		err = json.Unmarshal(buf.Bytes(), &data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v\n", data)
		// switch data["action"] {
		// case "Test":
		//	var t Test
		//	err = json.Unmarshal(buf.Bytes(), &t)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	fmt.Printf("%#v\n", t)
		// case "start":
		//	var s Start
		//	err = json.Unmarshal(buf.Bytes(), &s)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	fmt.Printf("%#v\n", s)
		// case "setFBTransparent":
		//	var a SetFBTransparent
		//	err = json.Unmarshal(buf.Bytes(), &a)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	fmt.Printf("%#v\n", a)
		// case "avSetLayout":
		//	var l AvLayout
		//	err = json.Unmarshal(buf.Bytes(), &l)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	fmt.Printf("%#v\n", l)
		// default:
		//	log.Fatal(err)
		// }
	}

}

func main() {
	TestService()
}
