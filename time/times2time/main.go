package main

import (
	"fmt"
	"time"
)

func main() {
	var ts int64
	ts = 1605228323
	tm := time.Unix(ts, 0)
	fmt.Println(tm.Format("2006-01-02"))
	date, _ := time.Parse("2006-01-02", tm.Format("2006-01-02"))
	fmt.Println(date)
}
