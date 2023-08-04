package main

import (
	"fmt"
	"sort"
	"time"
)

type reviews_data struct {
	review_id   string
	date        time.Time
	score       int
	firstname   string
	anonymous   bool
	review_text string
	title_text  string
	rating      float64
	upcount     int
}

type timeSlice []reviews_data

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	return p[i].date.Before(p[j].date)
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var reviews_data_map = make(map[string]reviews_data)
	reviews_data_map["1"] = reviews_data{date: time.Now().Add(12 * time.Hour)}
	reviews_data_map["2"] = reviews_data{date: time.Now()}
	reviews_data_map["3"] = reviews_data{date: time.Now().Add(24 * time.Hour)}
	//Sort the map by date
	date_sorted_reviews := make(timeSlice, 0, len(reviews_data_map))
	for _, d := range reviews_data_map {
		date_sorted_reviews = append(date_sorted_reviews, d)
	}
	fmt.Println(date_sorted_reviews)
	sort.Sort(date_sorted_reviews)
	fmt.Println(date_sorted_reviews)
}
