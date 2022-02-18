package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
)

func decode() {
	UserKey := "cef22d1b7d6e3ccf4aee5854f9af1767"
	UserKey1, _ := hex.DecodeString(UserKey)
	log.Printf("UserKey1:%s", UserKey1)
}

func main() {
	// MD5算法加密, 不可逆
	passwd := "passwd01"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(passwd))
	cipherKey := md5Ctx.Sum(nil)
	UserKey := hex.EncodeToString(cipherKey)
	log.Printf("UserKey:%s", UserKey) // UserKey:cef22d1b7d6e3ccf4aee5854f9af1767
	decode()
}
