package main

import (
	"encoding/json"
	"fmt"
)

func Data2Json(anyData interface{}) string {
	JsonByte, err := json.Marshal(anyData)
	if err != nil {
		fmt.Printf("解析戳:%s\n", err.Error())
	}
	return string(JsonByte)
}

func main() {
	data := []map[string]map[string]string{
		{"master": {"labels": "app:test"}},
	}
	//
	//for index, maps:=range data{
	//	if
	//}

	data[0]["master1"] = map[string]string{"a2": "2"}
	data[0]["master1"] = map[string]string{"a2": "2222"}

	fmt.Printf("data:\n%s\n", Data2Json(data))

}
