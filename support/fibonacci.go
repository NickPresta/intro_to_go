package main

import "fmt"

// START OMIT
func fibonacci(n int, ch chan uint64) {
	x, y := uint64(0), uint64(1)
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	ch := make(chan uint64, 42)
	go fibonacci(cap(ch), ch)
	fmt.Println("Starting...")
	for i := range ch {
		fmt.Println(i)
	}
}
// STOP OMIT
