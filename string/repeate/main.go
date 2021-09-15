// Go program to illustrate how to repeat
// a string to a specific number of times
package main

import (
	"fmt"
	"strings"
)

// Main method
func main() {

	// Creating and initializing a string
	// Using shorthand declaration
	str1 := "Welcome to GeeksforGeeks !.."
	str2 := "This is the tutorial of Go"

	// Repeating the given strings
	// Using Repeat function
	res1 := strings.Repeat(str1, 4)
	res2 := str2 + strings.Repeat("Language..", 2)

	// Display the results
	fmt.Println("Result 1: ", res1)
	fmt.Println("Result 2:", res2)
}
