package main

import "log"

// 截取子字符串
func Substr(input string, start int, length int, reverse bool) string {
	asRunes := []rune(input)
	if reverse {
		if len(asRunes) > length {
			start = len(asRunes) - length
		} else {
			start = 0
		}
		return string(asRunes[start:])
	} else {
		if start >= len(asRunes) {
			return ""
		}
		// 长度最长取字符全长
		if start+length > len(asRunes) {
			length = len(asRunes) - start
		}
		return string(asRunes[start : start+length])
	}
}

func main() {
	ret := Substr("helloewjglkwjgjgklwjeg", 0, 11, true)
	log.Printf("截断后:%+v", ret)
}
