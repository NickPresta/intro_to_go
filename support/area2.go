package main

import "fmt"

//define a Rectangle struct that has a length and a width
type Rectangle struct {
	length, width int
}

//write a function Area that can apply to a Rectangle type
func (r Rectangle) Area() int {
	return r.length * r.width
}

// START OMIT
type Shaper interface {
	Area() int
}

type Square struct {
	side int
}

func (sq Square) Area() int {
	return sq.side * sq.side
}

func main() {
	r := Rectangle{length: 5, width: 3}
	sq := Square{side: 4}
	shapes := []Shaper{r, sq}

	for i, _ := range shapes {
		fmt.Printf("Shape details: %+v\n", shapes[i])
		fmt.Println("Area: ", shapes[i].Area())
	}
}
// STOP OMIT
