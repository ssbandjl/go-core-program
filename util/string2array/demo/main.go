package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Prints:
// <nil> [[0 1] [0 2] [0 3]]
// 0 0 0
// 0 1 1
// 1 0 0
// 1 1 2
// 2 0 0
// 2 1 3

func main() {
	jsonstring := "[{\"name\":\"nginx\",\"version\":\"0.1.0\",\"description\":\"A Helm chart for Kubernetes\",\"apiVersion\":\"v1\",\"appVersion\":\"1.0\",\"urls\":[\"charts/nginx-0.1.0.tgz\"],\"created\":\"2021-03-04T09:57:46.924767435+08:00\",\"digest\":\"3828a4c60feff7971eeb10c9b5d4cc655a3fd6340fd6aa3e84a7de90b6f0655e\"}]"
	var listoflists []map[string]interface{}
	dec := json.NewDecoder(strings.NewReader(jsonstring))
	err := dec.Decode(&listoflists)
	fmt.Println(err, listoflists)
	//for i, list := range listoflists {
	//	for j, value := range list {
	//		fmt.Println(i, j, value)
	//	}
	//}
}