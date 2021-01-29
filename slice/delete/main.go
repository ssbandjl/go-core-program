package main

import "fmt"

func main() {
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(strSlice)

	strSlice = RemoveIndex(strSlice, 0)
	fmt.Println(strSlice)
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
