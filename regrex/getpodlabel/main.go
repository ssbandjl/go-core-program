package main

import (
	"fmt"
	"regexp"
)

func main() {

	// match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	// fmt.Println(match)

	str := "metadata.labels['stolon-cluster']"

	regexp, err := regexp.Compile(`\['(.*)'\]`)

	match := regexp.FindStringSubmatch(str)

	fmt.Println("Match: ", match[1], " Error: ", err)

	// str := "metadata.labels['stolon-cluster']"

	// regexp, _ := regexp.Compile("metadata\\.labels\\['(.*)'\\]")

	// fmt.Println(regexp.FindString(str))

	// match, _ := regexp.MatchString("metadata.labels[(.*)]", "metadata.labels['stolon-cluster']")
	// fmt.Println(match)

	// 	r, _ := regexp.Compile("p([a-z]+)ch")

	// 	fmt.Println(r.MatchString("peach"))

	// 	fmt.Println(r.FindString("peach punch"))

	// 	fmt.Println(r.FindStringIndex("peach punch"))

	// 	fmt.Println(r.FindStringSubmatch("peach punch"))

	// 	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// 	fmt.Println(r.FindAllStringSubmatchIndex(
	// 		"peach punch pinch", -1))

	// 	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 	fmt.Println(r.Match([]byte("peach")))

	// 	r = regexp.MustCompile("p([a-z]+)ch")
	// 	fmt.Println(r)

	// 	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// 	in := []byte("a peach")
	// 	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	// 	fmt.Println(string(out))

}
