package main

import "fmt"

// START OMIT
func change(slice []int) {
	slice[0] = -1
}

func main() {
	var evens []int
	evens = append(evens, 2, 4, 6, 8, 10)
	odds := []int{1, 3, 5, 7, 9}
	fmt.Println(odds, evens)

	change(evens)
	fmt.Println(evens)
}
// STOP OMIT
