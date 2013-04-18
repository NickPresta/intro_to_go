package main

import ("fmt";"time")

// START OMIT
func produce(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 42
}

func main() {
	ch := make(chan int)
	go produce(ch)
	fmt.Println("Waiting...")
	fmt.Println(<-ch)
}
// STOP OMIT
