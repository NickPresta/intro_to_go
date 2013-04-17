package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	later, _ := time.ParseDuration("5h")
	fmt.Printf("Later: %v\n", time.Now().Add(later))
}

// STOP OMIT
