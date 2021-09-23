// Go program to illustrate how to
// replace string with the specified regexp
package main

import (
	"fmt"
	"regexp"
)

// Main function
func main() {

	// Replace string with the specified regexp
	// Using ReplaceAllString() method
	m1 := regexp.MustCompile(`\s`)

	fmt.Println(m1.ReplaceAllString(`cloud1 688 
	 `, ""))
}
