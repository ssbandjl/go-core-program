package main

import "log"

func main() {
	log.Printf("\\n的字节数组:%+v", []byte("\n"))
	log.Printf("\\r的字节数组:%+v", []byte("\r"))
	log.Printf("\\r\\n的字节数组:%+v", []byte("\r\n"))
	log.Printf("一键三连:%+v", []byte("一键三连"))
	log.Printf("字节%+v", []byte("EOF"))
}
