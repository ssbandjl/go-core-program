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

/*现在我们有了适当大小的密钥，我们可以开始加密过程。 我们可以加密文本或任何二进制数据，这并不重要。在 Go 项目中，包含以下函数：*/
func encrypt(data []byte, passphrase string) []byte {
	// First we create a new block cipher based on the hashed passphrase. Once we have our block cipher, we want to wrap it in Galois Counter Mode (GCM) with a standard nonce length.
	/*NewGCM 返回给定的 128 位分组密码，以标准随机数长度封装在伽罗瓦计数器模式中。一般来说，这种 GCM 实现所执行的 GHASH 操作不是恒定时间的。 一个例外是当底层块由 aes.NewCipher 在硬件支持 AES 的系统上创建时。 有关详细信息，请参阅 crypto/aes 包文档。*/
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	// Before we can create the ciphertext, we need to create a nonce. 创建随机数
	/*我们创建的 nonce 需要是 GCM 指定的长度。 需要注意的是，用于解密的 nonce 必须与用于加密的 nonce 相同。有一些策略可用于确保我们的解密 nonce 与加密 nonce 匹配。 如果要进入数据库，一种策略是将随机数与加密数据一起存储。 另一种选择是将随机数预先或附加到加密数据中。 我们将添加随机数。*/
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	// Seal 命令中的第一个参数是我们的前缀值。 加密的数据将附加到它上面。 使用密文，我们可以将其返回给调用函数。
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
