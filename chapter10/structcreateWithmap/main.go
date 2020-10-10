package main

import "fmt"

type SvcInfo struct {
	kubernetesAlias string
	kubernetesNs    string
	kubernetesSvc   string
	instanceType    string
	labels          map[string]string
}

func main() {
	svcInfo := SvcInfo{
		kubernetesSvc: "ddd",
	}
	fmt.Printf("%+v", svcInfo)
}
