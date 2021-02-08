package main

import "log"

func main() {

	str := "123456789"

	strLen := len(str)
	log.Printf("str length:%d", strLen)
	if strLen >= 3 {
		str = str[strLen-3:]
	}
	log.Printf("str cut result:%s", str)
}
