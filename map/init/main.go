package main

import "fmt"

var (
	PrometheusClientCache map[string]string
)

func init() {
	if PrometheusClientCache == nil {
		fmt.Println("init")
		PrometheusClientCache = make(map[string]string)
	}
}

func main() {
	PrometheusClientCache["192.168.1.10"] = "xxxxxx"
	fmt.Println(PrometheusClientCache)
}
