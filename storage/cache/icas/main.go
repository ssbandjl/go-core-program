package main

import (
	"bytes"
	"log"
	"os/exec"
)

func GetNtpSyncStatus() bool {
	var stdOut, stdErr bytes.Buffer
	c := "casadm -L -o csv"
	cmd := exec.Command("/bin/bash", "-c", c)
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		log.Printf("ntpq -c rv error, stdout:%s, stderr:%s", stdOut.String(), stdErr.String())
		return false
	}
	return true
}

func main() {}
