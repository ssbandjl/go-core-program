package main

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

func main() {
	// loop
	for {
		var (
			s string
		)
		log.Printf("请输入IP地址:\n")
		_, err := fmt.Scan(&s)
		if err == io.EOF {
			log.Printf("io.EOF:%s", err.Error())
			return
		}
		ips := strings.Split(s, ".")
		if len(ips) != 4 { // IP,区段数量不够
			fmt.Println("NO")
			return
		}
		var flag bool
		for i := range ips {
			num, err := strconv.Atoi(ips[i]) // 字符串转int
			if err != nil {
				fmt.Println("NO")
				return
			}
			if num < 0 || num > 255 { // 比较每个区段大小
				flag = true
				break
			}
		}
		if !flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
