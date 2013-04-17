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
	go say("ping")
	fmt.Println("pong")
	time.Sleep(3 * time.Second)
}
// STOP OMIT
