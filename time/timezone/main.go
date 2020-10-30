package main

import (
	"fmt"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

func parseWithLocation(name string, timeStr string) (time.Time, error) {
	locationName := name
	if l, err := time.LoadLocation(locationName); err != nil {
		println(err.Error())
		return time.Time{}, err
	} else {
		lt, _ := time.ParseInLocation(TIME_LAYOUT, timeStr, l)
		fmt.Println(locationName, lt)
		return lt, nil
	}
}
func main() {
	fmt.Println("0. now: ", time.Now())
	str := "2018-09-10 00:00:00"
	fmt.Println("1. str: ", str)

	//字符串转时间
	t, _ := time.Parse(TIME_LAYOUT, str)
	fmt.Println("2. Parse time: ", t)

	tStr := t.Format(TIME_LAYOUT)
	fmt.Println("3. Format time str: ", tStr)
	name, offset := t.Zone()
	name2, offset2 := t.Local().Zone()

	fmt.Printf("4. Zone name: %v, Zone offset: %v\n", name, offset)
	fmt.Printf("5. Local Zone name: %v, Local Zone offset: %v\n", name2, offset2)
	tLocal := t.Local()
	tUTC := t.UTC()
	fmt.Printf("6. t: %v, Local: %v, UTC: %v\n", t, tLocal, tUTC)
	fmt.Printf("7. t: %v, Local: %v, UTC: %v\n", t.Format(TIME_LAYOUT), tLocal.Format(TIME_LAYOUT), tUTC.Format(TIME_LAYOUT))
	fmt.Printf("8. Local.Unix: %v, UTC.Unix: %v\n", tLocal.Unix(), tUTC.Unix())
	str2 := "1969-12-31 23:59:59"
	t2, _ := time.Parse(TIME_LAYOUT, str2)
	fmt.Printf("9. str2：%v，time: %v, Unix: %v\n", str2, t2, t2.Unix())
	fmt.Printf("10. %v, %v\n", tLocal.Format(time.ANSIC), tUTC.Format(time.ANSIC))
	fmt.Printf("11. %v, %v\n", tLocal.Format(time.RFC822), tUTC.Format(time.RFC822))
	fmt.Printf("12. %v, %v\n", tLocal.Format(time.RFC822Z), tUTC.Format(time.RFC822Z))

	//指定时区
	parseWithLocation("America/Cordoba", str)
	parseWithLocation("Asia/Shanghai", str)
	parseWithLocation("Asia/Beijing", str)
}
