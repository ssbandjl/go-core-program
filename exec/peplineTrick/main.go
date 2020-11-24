package main

//上面的解决方案是Go风格的解决方案，事实上你还可以用一个"Trick"花招/hack方式/来实现。

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := "cat /proc/cpuinfo | egrep '^model name' | uniq | awk '{print substr($0, index($0,$4))}'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	fmt.Println(string(out))
}
