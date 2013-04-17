package main

import "fmt"

// START OMIT
type ProperName interface {
	FullName() string
}

type Dog struct {
	FirstName string
}

func main() {
	var proper ProperName
	p := Person{"John", "Smith"}
	d := Dog{"Lassie"}

	proper = p
	proper = d

	fmt.Println(proper.FullName())
}
// STOP OMIT

type Person struct {
	FirstName string
	LastName string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}
