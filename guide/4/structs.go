package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X:1}
	v3 = Vertex{}
	vp = &Vertex{1, 2}
)
func main() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	p := &v
	p.X = 1e9
	fmt.Println(v.X)

	fmt.Println(v1, vp, v2, v3)
}
