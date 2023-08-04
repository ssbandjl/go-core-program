package main

import (
	"bytes"
	"log"
	"os/exec"
)

func GetNtpSyncStatus() bool {
	var out bytes.Buffer
	c := "ntpq -c rv|grep leap_none"
	cmd := exec.Command("/bin/bash", "-c", c)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Printf("error:%s", err.Error())
		return false
	}
	log.Printf("out:%s", out.String())
	return true
}

func main() {
	if GetNtpSyncStatus() {
		log.Printf("同步成功")
	} else {
		log.Printf("同步失败")
	}
}
