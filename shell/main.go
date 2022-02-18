package main

import "github.com/codeskyblue/go-sh"

func main() {
	sh.Command("echo", "hello\tworld").Command("cut", "-f2").Run()

	session := sh.NewSession()
	session.ShowCMD = true
	session.Command("echo", "hello").Run()
	// set ShowCMD to true for easily debug
}
