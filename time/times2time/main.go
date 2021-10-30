package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	var ts int64
	ts = 1605228323
	tm := time.Unix(ts, 0)
	fmt.Println(tm.Format("2006-01-02"))
	date, _ := time.Parse("2006-01-02", tm.Format("2006-01-02"))
	fmt.Println("date", date)

	timeStr := "210102 13:14:15"
	log.Printf("时间格式化:%s", tm.Format("200102 15:04:05"))
	timeDate, err := time.Parse(timeStr, "200102 15:04:05")
	if err != nil {
		log.Printf("时间解析出错:%s", err.Error())
	}
	log.Printf("timeDate:%v", timeDate)
	fmt.Println(time.Now().Format("02-Jan-2006 15:04:05"))
}
