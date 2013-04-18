package main

import (
	"fmt"
	"time"
)

// START OMIT
func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("ping") // Starts immediately, doesn't block
	fmt.Println("pong") // continues
	time.Sleep(3 * time.Second)
}
// STOP OMIT
