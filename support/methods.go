package main

import "fmt"

// START OMIT
type Person struct {
	FirstName string
	LastName string
}

func (p *Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func main() {
	p := Person{"John", "Smith"}
	fmt.Println(p.FullName())
}
// STOP OMIT
