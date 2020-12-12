package main

import (
	"log"
)

func main() {
	var jobName string
	strLong := "rdm-mongodb-headless-harix-oma-roc-cloudpepper-v3-fit"
	if len(strLong) >= 52 {
		jobName = strLong[0:51]
		log.Printf("%s", jobName)
		if jobName[50:51] == "f" {
			jobName = jobName[0:50]
			log.Printf("%s", jobName)
		}
	}
}
