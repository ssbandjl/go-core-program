package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// go run main.go
func main() {
	result, err := ping("8.8.8.8")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(result)

}

func ping(ip string) (result string, err error) {
	var out, stdErr bytes.Buffer
	c := fmt.Sprintf("./ping %s -c10", ip)
	cmd := exec.Command("sh", "-c", c)
	cmd.Stdout = &out
	cmd.Stderr = &stdErr
	err = cmd.Run()
	if err != nil {
		fmt.Println(stdErr.String())
		fmt.Println(err.Error())
		return "", err
	}
	// fmt.Println(out.String())
	return out.String(), nil
}
