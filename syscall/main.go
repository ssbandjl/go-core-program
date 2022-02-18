package main

import (
	"log"
	"syscall"
)

func main() {
	log.Printf("syscall.ETIMEDOUT:%v", syscall.ETIMEDOUT)
}
