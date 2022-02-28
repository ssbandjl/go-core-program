package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/forgoer/openssl"
)

func main() {
	src := []byte("passwd01")
	// key := []byte("1234567890123456")
	key := []byte("NeonSanEncrypted")
	dst, err := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	if err != nil {
		log.Printf("加密错误:%s", err.Error())
	}
	// fmt.Println("加密后:", dst)
	encrypt_str := base64.StdEncoding.EncodeToString(dst)
	log.Printf("加密后:%s", encrypt_str)
	// fmt.Printf(base64.StdEncoding.EncodeToString(dst)) // yXVUkR45PFz0UfpbDB8/ew==
	dst, _ = openssl.AesECBDecrypt(dst, key, openssl.PKCS7_PADDING)
	fmt.Println(string(dst)) // 123456
}
