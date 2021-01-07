package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("Yesterday: %v\n", time.Now().Add(-24*time.Hour))

	fmt.Printf("%+v\n", time.Now().Local().UnixNano())
	fmt.Printf("当前时间戳:%+v\n", time.Now().Unix()+10)
	fmt.Printf("当前时间:%+v\n", time.Now())

	//
	//loc, _ := time.LoadLocation("Asia/Shanghai")
	loc, _ := time.LoadLocation("UTC")
	fmt.Printf("时区:%s, 时间:%s\n", loc, time.Now().In(loc).Format("2006-01-02 15:04:05"))

	//30天前
	someDays := 30
	//beforeSomeDays := time.Now().Add(-someDays*24*time.Hour)
	beforeSomeDays := time.Now().AddDate(0, 0, -someDays)

	fmt.Printf("30天前: %s", beforeSomeDays.Format("2006-01-02"))

}
