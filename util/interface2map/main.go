package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	b := []byte(`{"key":"value"}`)

	var f interface{}
	json.Unmarshal(b, &f)

	myMap := f.(map[string]interface{})

	fmt.Println(myMap["key"])
}