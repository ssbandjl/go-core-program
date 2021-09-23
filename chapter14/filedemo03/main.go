package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//使用ioutil.ReadFile一次性将文件读取到位
	// file := "d:/test.txt"
	file := "/Users/xb/Downloads/often/tmp/tmp.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	//把读取到的内容显示到终端
	//fmt.Printf("%v", content) // []byte
	fmt.Printf("内容:%v\n", string(content))              // []byte
	fmt.Printf("字节长度:%v\n", len(string(content)))       // []byte
	fmt.Printf("长度:%v\n", len([]rune(string(content)))) // []byte
	fmt.Printf("%v", string(content)[0:65535])

	//我们没有显式的Open文件，因此也不需要显式的Close文件
	//因为，文件的Open和Close被封装到 ReadFile 函数内部
}
