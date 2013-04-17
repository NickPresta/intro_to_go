package main

import "fmt"

// START OMIT
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{Y:3})

	v := Vertex{X:1, Y:2}
	v.X = 3
	fmt.Println(v)

	p := &v
	p.X = 6
	fmt.Println(v)
}
// STOP OMIT
