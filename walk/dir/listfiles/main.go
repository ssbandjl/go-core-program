package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

func main(){
	files, err := ioutil.ReadDir("/Users/xb/Downloads/wordpress2/")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
