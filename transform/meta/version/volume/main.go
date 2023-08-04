package main

import "log"

// constant about volume info
const (
	MaxMetaVer = 0x7fff
)

func main() {
	log.Printf("%d", (MaxMetaVer&0xffff)<<48) // 9223090561878065152
}
