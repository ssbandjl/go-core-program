package main

import (
	"fmt"
	"math/big"
)

func main() {

	// oldNum := float64(8.0497183772403904e+17)
	oldNum := float64(6.724456448e+09)
	newNum := big.NewRat(1, 1)
	newNum.SetFloat64(oldNum)
	fmt.Println(newNum.FloatString(0))
}
