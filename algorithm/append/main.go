package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var num int
var output []string

func main() {

	var input [][]string

	//     fmt.Scanln(&input)
	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	for scanner.Scan() {
		//         fmt.Println(scanner.Text())
		if i == 0 {
			num, _ = strconv.Atoi(scanner.Text())
		} else {
			arr := strings.Split(scanner.Text(), ",")
			input = append(input, arr)
		}
		i++
	}
	run(input)
	fmt.Println(strings.Join(output, ","))
	fmt.Println(input)

}

func run(input [][]string) {
	if len(input) == 0 {
		return
	}
	for i, item := range input {
		if len(item) >= num {
			for j := 0; j < num; j++ {
				output = append(output, item[j])
				input[i] = item[(len(item) - num - 1):]
			}

		} else {
			for j := 0; j < len(item); j++ {
				output = append(output, item[j])
				input[i] = item[:0]
			}
		}
	}
}
