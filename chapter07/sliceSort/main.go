package main

import (
	"log"
	"sort"
)

func main() {
	slice := []float64{111.0, 222.111, 333.222, 0, 100.111, 50.224}
	sort.Float64s(slice)
	log.Printf("sorted:%+v", slice)

}
