package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type ShoppingRecord struct {
	// 1. Create a struct for storing CSV lines and annotate it with JSON struct field tags
	Vegetable string `json:"vegetable"`
	Fruit     string `json:"fruit"`
	Rank      int    `json:"rank"`
}

func createShoppingList(data [][]string) []ShoppingRecord {
	// convert csv lines to array of structs
	log.Printf("data:\n%s", data)
	var shoppingList []ShoppingRecord
	for i, line := range data { // index, line
		if i > 0 { // omit header line
			var rec ShoppingRecord
			for j, field := range line {
				if j == 0 {
					rec.Vegetable = field
				} else if j == 1 {
					rec.Fruit = field
				} else if j == 2 {
					var err error
					rec.Rank, err = strconv.Atoi(field)
					if err != nil {
						continue
					}
				}
			}
			shoppingList = append(shoppingList, rec)
		}
	}
	return shoppingList
}

func main() {
	// open file
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// 2. Read CSV file using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 3. Assign successive lines of raw CSV data to fields of the created structs
	shoppingList := createShoppingList(data)

	// 4. Convert an array of structs to JSON using marshaling functions from the encoding/json package
	// MarshalIndent 类似于 Marshal，但应用 Indent 来格式化输出。 根据缩进嵌套，输出中的每个 JSON 元素都将在以前缀开头的新行开始，后跟一个或多个缩进副本。
	jsonData, err := json.MarshalIndent(shoppingList, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}
