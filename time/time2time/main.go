package main

import (
	"log"
	"strings"
	"time"

	"github.com/viant/toolbox"
)

func main() {
	log.Printf("时间转换")
	// dateLaout := toolbox.DateFormatToLayout("yyyy-MM-dd hh:mm:ss z")
	dateLaout := toolbox.DateFormatToLayout("yyMMdd")
	timeValue, err := time.Parse(dateLaout, "210222")
	// timeValue, err := time.Parse(dateLaout, "210125 20:19:10 UTC")
	ymd := strings.Split(timeValue.String(), " ")
	if err != nil {
		log.Printf("解析出错:%s", err.Error())
	}
	log.Printf("结果:%v, 年月日:%s", timeValue.String(), ymd[0])
}
