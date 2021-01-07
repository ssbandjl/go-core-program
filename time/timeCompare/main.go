package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now()
	tomorrow := today.Add(24 * time.Hour)
	sameday := tomorrow.Add(-24 * time.Hour)
	someDayBefore := today.AddDate(0, 0, -30)

	b1 := today.Before(tomorrow)
	fmt.Println("today before tomorrow:", b1)

	b2 := tomorrow.After(today)
	fmt.Println("tomorrow after today:", b2)

	if today != tomorrow {
		fmt.Println("today is not tomorrow")
	}

	if sameday == today {
		fmt.Println("sameday is today")
	}

	if today.Equal(sameday) {
		fmt.Println("today is sameday")
	}
	fmt.Println(today.After(someDayBefore))
}
