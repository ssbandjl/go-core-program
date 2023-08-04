package main

import (
	"fmt"
	"math"
	"strconv"
)

// KeepDecimal keep decimal
func KeepDecimal(floatInput float64, num float64) float64 {
	return math.Trunc(floatInput*math.Pow(10, num)) / math.Pow(10, num)
}

// 23949811384320/23992761057280
func main() {
	var a uint64 = 23949811384320
	var b uint64 = 23992761057280
	fmt.Println(1.0 - float64(a)/float64(b))
	str := fmt.Sprintf("%.2f", (1.0-float64(a)/float64(b))*100)
	float64Num, _ := strconv.ParseFloat(str, 64)
	fmt.Println(float64Num)
	fmt.Println(KeepDecimal(0.091111888999, -1))
}
