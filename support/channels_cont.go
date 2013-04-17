package main

import "fmt"

// START OMIT
func main() {
	c := make(chan int)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
// STOP OMIT
