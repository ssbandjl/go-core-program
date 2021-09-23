package main

import "log"

// 截取子字符串
func Substr(input string, start int, length int, reverse bool) string {
	// asRunes := []rune(input)
	if reverse {
		if len(input) > length {
			start = len(input) - length
		} else {
			start = 0
		}
		return string(input[start:])
	} else {
		if start >= len(input) {
			return ""
		}
		// 长度最长取字符全长
		if start+length > len(input) {
			length = len(input) - start
		}
		return input[start : start+length]
	}
}

func main() {
	ret := Substr("中国你好", 0, 3, true)
	log.Printf("截断后:%+v, 字节:%v, 长度:%v", ret, len(ret), len([]rune(ret)))
}
