package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	//base64加密
	passwordStr := "root"
	encodeStr := b64.StdEncoding.EncodeToString([]byte(passwordStr))
	decodeStr, _ := b64.StdEncoding.DecodeString(encodeStr)
	fmt.Printf("加密前:%s, 加密后:%s, 解密后:%s", passwordStr, encodeStr, decodeStr)
}
