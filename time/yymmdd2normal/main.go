package main

import (
	"errors"
	"log"
	"strings"
)

func Yymmdd2Normal(dateTime string) (string, error) {
	dateTimeSplit := strings.Split(dateTime, " ") //[210130 21:45:02]
	if len(dateTimeSplit[0]) != 6 {
		return "", errors.New("date time length error")
	}
	//
	dateByte := []byte(dateTimeSplit[0])
	dateByte20 := append([]byte("20"), dateByte...)
	dateByteY := append(dateByte20[:4], []byte("-")...)
	dateByteM := append(dateByteY, dateByte[2:4]...)
	dateByteM = append(dateByteM, []byte("-")...)
	dateByteD := append(dateByteM, dateByte[4:6]...)
	// log.Printf("dateByteD:%s", dateByteD)
	// log.Printf("dateByte:%v, yymmdd:%s", dateByte, dateByte)
	result := string(dateByteD) + " " + dateTimeSplit[1]
	log.Printf("result:%s", result)
	return result, nil
}

func main() {
	Yymmdd2Normal("210130 21:45:02")
}
