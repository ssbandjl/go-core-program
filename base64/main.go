package main

import (
	b64 "encoding/base64"
	"fmt"
	"strings"
)

func main() {
	//base64加密

	//输入待编码的密码
	passwordStr := "cloud1688"
	encodeStr := b64.StdEncoding.EncodeToString([]byte(passwordStr))
	decodeStr, _ := b64.StdEncoding.DecodeString(encodeStr)
	fmt.Printf("加密前:%s, 加密后:%s, 解密后:%s\n\n", passwordStr, encodeStr, decodeStr)
	fmt.Printf("加密前:%s, 加密后:%v, 解密后:%s\n\n", passwordStr, []byte(encodeStr), decodeStr)

	//输入待解码的密码
	DecodeStr := "Y2xvdWQxNjg4CiA="
	DecodeStrResult, _ := b64.StdEncoding.DecodeString(DecodeStr)

	// 替换
	DecodeStrResultTmp := strings.ReplaceAll(string(DecodeStrResult), "\n", "")
	DecodeStrResultTmp = strings.ReplaceAll(DecodeStrResultTmp, "\r", "")
	fmt.Printf("Base64解密前:%s\nBase64解密后:%sEND\n", DecodeStr, DecodeStrResultTmp)
	fmt.Printf("Base64解密前:%s\nBase64解密后字节:%vEND\n", DecodeStr, []byte(DecodeStrResultTmp))
}
