package main

// 参考:Encrypt And Decrypt Data In A Golang Application With The Crypto Packages: https://www.thepolyglotdeveloper.com/2018/02/encrypt-decrypt-data-golang-application-crypto-packages/

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

// 解密
func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	log.Printf("decrypt: hash key:%s", key)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func encryptFile(filename string, data []byte, passphrase string) {
	f, _ := os.Create(filename)
	defer f.Close()
	f.Write(encrypt(data, passphrase))
}

func decryptFile(filename string, passphrase string) []byte {
	data, _ := ioutil.ReadFile(filename)
	return decrypt(data, passphrase)
}

func main() {
	// 开始加密
	ciphertext := encrypt([]byte("Hello World"), "password")
	fmt.Printf("Encrypted: %x\n", ciphertext) // 1c9da161adc0b28164e600c3ad37a0fc4f4f4a4a6592e080e3e31b800aec7ec581135c9eb5fc17
	plaintext := decrypt(ciphertext, "password")
	fmt.Printf("Decrypted: %s\n", plaintext)
	// 加密到文件
	encryptFile("sample.txt", []byte("Hello World"), "password1")
	// 从文件中解密
	fmt.Println(string(decryptFile("sample.txt", "password1")))
}
