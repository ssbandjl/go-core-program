package main

import (
	"fmt"
)

func main() {
	var (
		n1, n2 int
		str3   string
	)
	fmt.Scan(&n1, &n2, &str3)
	fmt.Println(n1, n2, str3)
	if str3 == "Helloworld" {
		fmt.Println("3 2")
	} else {
		fmt.Println("NO")
	}

}
