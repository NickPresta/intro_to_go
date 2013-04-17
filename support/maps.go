package main

import "fmt"

// START OMIT
func main() {
	mapping := map[string]string{"foo":"bar", "Jane":"Doe"}
	mapping["editor"] = "Vi"
	fmt.Println(mapping)
	mapping["foo"] = "baz"
	fmt.Println(mapping["foo"])
	delete(mapping, "Jane")
	fmt.Println(mapping)

	// Check for membership
	value, ok := mapping["nick"]
	if ok {
		fmt.Printf("Value = '%s'\n", value)
	} else {
		fmt.Println("nick not in map!")
	}
}
// STOP OMIT
