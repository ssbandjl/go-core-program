package main

import (
	"fmt"
)

// lib1.go in p3

var V1_p3 = Trace("init v1_p3", 3)
var V2_p3 = Trace("init v2_p3", 3)

func init() {
	fmt.Println("init func in p3")
	V1_p3 = 300
	V2_p3 = 300
}
