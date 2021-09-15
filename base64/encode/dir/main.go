package main

import (
	"bufio"
	"encoding/base64"
	"io/ioutil"
	"log"
	"fmt"
	"os"
)


func main(){
	log.Printf("将目录编码为文件")
	// Open file on disk.
	f, _ := os.Open("/Users/xb/Downloads/wordpress2/Chart.yaml")

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	// Print encoded data to console.
	// ... The base64 image can be used as a data URI in a browser.
	fmt.Println("ENCODED: " + encoded)
}
