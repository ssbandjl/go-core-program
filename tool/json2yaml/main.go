package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"github.com/ghodss/yaml"
	"os"
)

func main() {
	//使用ioutil.ReadFile一次性将文件读取到位
	file := "input.json"
	j, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}
	//把读取到的内容显示到终端
	//fmt.Printf("%v", content) // []byte
	//fmt.Printf("%v", string(content)) // []byte


	//j := []byte(`{"name": "John", "age": 30}`)
	y, err := yaml.JSONToYAML(j)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))


	filePath := "output.yaml"
	fileOutput, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//及时关闭file句柄
	defer fileOutput.Close()

	writer := bufio.NewWriter(fileOutput)
	writer.WriteString(string(y))
	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer.Flush()

	/* Output:
	name: John
	age: 30
	*/
	//j2, err := yaml.YAMLToJSON(y)
	//if err != nil {
	//	fmt.Printf("err: %v\n", err)
	//	return
	//}
	//fmt.Println(string(j2))
	/* Output:
	{"age":30,"name":"John"}
	*/

	// go run main.go
}