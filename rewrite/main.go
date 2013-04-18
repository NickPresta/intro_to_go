package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	all := numbers[0:len(numbers)]
	fmt.Println(all)
}
