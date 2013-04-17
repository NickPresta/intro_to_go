package main

import "fmt"

// START OMIT
func main() {
	var answer int
	answer = 42

	greeting := "Hello"
	awesome := true
	var maxInt uint64 = 1<<64 - 1
	complexNum := 1+0i
	var zeroed []int

	const f = "%T(%v)\n"
	fmt.Printf(f, answer, answer)
	fmt.Printf(f, greeting, greeting)
	fmt.Printf(f, awesome, awesome)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, complexNum, complexNum)
	fmt.Printf(f, zeroed, zeroed)
}
// STOP OMIT
