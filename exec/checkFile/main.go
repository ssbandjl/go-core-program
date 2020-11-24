package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
)

//想象一下你写了一个程序需要花费很长时间执行，再最后你调用foo做一些基本的任务。
//
//如果foo程序不存在，程序会执行失败。
//
//当然如果我们预先能检查程序是否存在就完美了，如果不存在就打印错误信息。
//
//你可以调用exec.LookPath方法来检查：


func checkLsExists() {
	path, err := exec.LookPath("ls")
	if err != nil {
		fmt.Printf("didn't find 'ls' executable\n")
	} else {
		fmt.Printf("'ls' executable is in '%s'\n", path)
	}
}

func main(){
	cmd := exec.Command("programToExecute")
	additionalEnv := "FOO=bar"
	newEnv := append(os.Environ(), additionalEnv))
	cmd.Env = newEnv
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("%s", out)
}

//另一个检查的办法就是让程序执行一个空操作， 比如传递参数"--help"显示帮助信息。
//
//下面的章节是译者补充的内容