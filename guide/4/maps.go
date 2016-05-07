package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex {
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var ma = map[string]Vertex {
		"Bell Labs": Vertex {
			40.123, -567.8,
		},
		"Google": Vertex{
			123.45, -648.12,
		},
	}
	fmt.Println(ma)

	var mma = map[string]Vertex{
		"Bell Labs" : {123.45, -163234.23},
		"Google" : {1233.14, -23.234},
	}
	fmt.Println(mma)


	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value: ", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value: ", v, "Present?", ok)
}
