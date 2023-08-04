package main

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/forgoer/openssl"
)

func main() {
	src := []byte("1567")
	// key := []byte("1234567890123456")
	key := []byte("NeonSanEncrypted")
	dst, err := openssl.AesECBEncrypt(src, key, openssl.PKCS7_PADDING)
	if err != nil {
		log.Printf("加密错误:%s", err.Error())
	}
	fmt.Println("加密后byte数组:", dst)
	encrypt_str := base64.StdEncoding.EncodeToString(dst)
	log.Printf("加密后:%s, 长度:%d", encrypt_str, len(encrypt_str))
	// fmt.Printf(base64.StdEncoding.EncodeToString(dst)) // yXVUkR45PFz0UfpbDB8/ew==
	base64Decode, err := base64.StdEncoding.DecodeString(encrypt_str)
	if err != nil {
		log.Printf("base64Decode err:%s", err.Error())
	}
	dst, err = openssl.AesECBDecrypt(base64Decode, key, openssl.PKCS7_PADDING)
	if err != nil {
		log.Printf("解密错误:%s", err.Error())
	}
	// dst, _ = openssl.AesECBDecrypt(dst, key, openssl.PKCS7_PADDING)
	fmt.Println("解密后:", string(dst)) // 123456
}
