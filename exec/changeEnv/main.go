package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

//改变执行程序的环境(environment)
//你已经知道了怎么在程序中获得环境变量，对吧： ｀os.Environ()｀返回所有的环境变量[]string,每个字符串以FOO=bar格式存在。FOO是环境变量的名称，bar是环境变量的值， 也就是os.Getenv("FOO")的返回值。
//
//有时候你可能想修改执行程序的环境。
//
//你可设置exec.Cmd的Env的值，和os.Environ()格式相同。通常你不会构造一个全新的环境，而是添加自己需要的环境变量：

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

//包 shurcooL/go/osutil提供了便利的方法设置环境变量。