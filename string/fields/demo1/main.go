package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	out := `
169.254.2.1 ens49f0
169.254.3.1 ens49f1
	`
	data_name2Gateway := make(map[string]string)
	spineLeafPort := strings.Fields(out)
	log.Printf("spineLeafPort:%s", spineLeafPort)
	for j := 0; j < len(spineLeafPort); j = j + 2 {
		data_name2Gateway[spineLeafPort[j+1]] = spineLeafPort[j]
	}
	fmt.Println(data_name2Gateway)
}
