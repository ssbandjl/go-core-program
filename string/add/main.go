package main

import "fmt"

func AddStr(data *string, addStr string) {
	*data = *data + addStr + "\r\n"
}

func main() {
	var data string
	data = ""
	addStr := "增加的数据"
	AddStr(&data, addStr)
	fmt.Printf("增加后的字符串:\n%s\n", data)
}
