package main

import (
	"fmt"
)

var (
	input  string
	myline []string
)

func main() {
	tmp1 := ""
	while (true){
		n, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(n)
		fmt.Println(input)
		if input != "\n" {
			tmp1 += input
		}else{
			continue
		}
	}

	fmt.Println(tmp1)
}
