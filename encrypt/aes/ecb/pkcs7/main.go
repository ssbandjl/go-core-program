// Golang里的AES加密、解密，支持AES-ECB-PKCS7Padding等多种加密组合，兼容JAVA、PHP等语言:https://tech1024.com/original/3015

package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
)

var (
	// ErrInvalidBlockSize indicates hash blocksize <= 0.
	ErrInvalidBlockSize = errors.New("invalid blocksize")

	// ErrInvalidPKCS7Data indicates bad input to PKCS7 pad or unpad.
	ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data (empty or not padded)")

	// ErrInvalidPKCS7Padding indicates PKCS7 unpad fails to bad input.
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

func PKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

func PKCS7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	// log.Printf("src len:%v", len(src))
	// log.Printf("len(src)%%blockSize:%v", len(src)%blockSize)
	// log.Printf("padding:%v", padding)
	// fmt.Println(byte(padding))
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	// log.Printf("padtext:%s", padtext)
	return append(src, padtext...)
}

func PKCS7UnPadding(src []byte) []byte {
	// length := len(src)
	// unpadding := int(src[length-1])
	// return src[:(length - unpadding)]
	length := len(src)
	if length == 0 {
		return src
	}
	unpadding := int(src[length-1])
	if length < unpadding {
		return src
	}
	return src[:(length - unpadding)]
}

func pkcs7Pad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	n := blocksize - (len(b) % blocksize)
	pb := make([]byte, len(b)+n)
	copy(pb, b)
	copy(pb[len(b):], bytes.Repeat([]byte{byte(n)}, n))
	return pb, nil
}

// pkcs7Unpad validates and unpads data from the given bytes slice.
// The returned value will be 1 to n bytes smaller depending on the
// amount of padding, where n is the block size.
func pkcs7Unpad(b []byte, blocksize int) ([]byte, error) {
	if blocksize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if b == nil || len(b) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	if len(b)%blocksize != 0 {
		return nil, ErrInvalidPKCS7Padding
	}
	c := b[len(b)-1]
	n := int(c)
	if n == 0 || n > len(b) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if b[len(b)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return b[:len(b)-n], nil
}

func ECBEncrypt(block cipher.Block, src, key []byte) ([]byte, error) {
	blockSize := block.BlockSize()

	encryptData := make([]byte, len(src))
	tmpData := make([]byte, blockSize)

	for index := 0; index < len(src); index += blockSize {
		block.Encrypt(tmpData, src[index:index+blockSize])
		copy(encryptData, tmpData)
	}
	return encryptData, nil
}

func ECBDecrypt(block cipher.Block, src, key []byte) ([]byte, error) {
	dst := make([]byte, len(src))

	blockSize := block.BlockSize()
	tmpData := make([]byte, blockSize)

	for index := 0; index < len(src); index += blockSize {
		block.Decrypt(tmpData, src[index:index+blockSize])
		copy(dst, tmpData)
	}

	return dst, nil
}

func main() {

	// 加密
	src := []byte("123456jswkjegklawjgekljkwg2kl34jq90235jklgj~~jklsjgk")
	key := []byte("1234567890123456") //16个字节的密钥

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//PKCS7Padding填充
	// log.Printf("block.BlockSize():%d", block.BlockSize())
	src = PKCS7Padding(src, block.BlockSize())

	// ECB加密
	dst, err := ECBEncrypt(block, src, key)
	if err != nil {
		panic(err)
	}

	fmt.Println("加密后:", base64.StdEncoding.EncodeToString(dst)) // SpfAShHImQhWjd/21Pgz2Q==

	// 解密
	src, err = ECBDecrypt(block, dst, key)
	if err != nil {
		panic(err)
	}

	src = PKCS7UnPadding(src)

	fmt.Println("解密后:", string(src)) // 123456
	log.Printf("解密后:%+v", string(src))

}
