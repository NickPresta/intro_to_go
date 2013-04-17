package main

import "fmt"


// START OMIT
func add(numbers ...int) (int, int) {
	total := 0
	for _, value := range numbers {
		total = total + value
	}
	return len(numbers), total
}

func main() {
	count, total := add(1, 3, 5, 7, 9)
	fmt.Printf("The total is %d from %d numbers\n", total, count)
}
// STOP OMIT
