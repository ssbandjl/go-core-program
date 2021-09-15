package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Prints:
// <nil> [[0 1] [0 2] [0 3]]
// 0 0 0
// 0 1 1
// 1 0 0
// 1 1 2
// 2 0 0
// 2 1 3

func main() {
	jsonstring := "[[0, 1], [0, 2], [0,3 ]]"
	var listoflists [][]int
	dec := json.NewDecoder(strings.NewReader(jsonstring))
	err := dec.Decode(&listoflists)
	fmt.Println(err, listoflists)
	for i, list := range listoflists {
		for j, value := range list {
			fmt.Println(i, j, value)
		}
	}
}