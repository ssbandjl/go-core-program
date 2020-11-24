//package pepline
//
//import (
//	"bytes"
//	"io"
//	"os"
//)
//
//我们可以使用管道将多个命令串联起来， 上一个命令的输出是下一个命令的输入。
//
//使用os.Exec有点麻烦，你可以使用下面的方法：

package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
)

func main() {
	c1 := exec.Command("ls")
	c2 := exec.Command("wc", "-l")
	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r
	var b2 bytes.Buffer
	c2.Stdout = &b2
	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
	io.Copy(os.Stdout, &b2)
}

//或者直接使用Cmd的StdoutPipe方法，而不是自己创建一个io.Pipe`。
//
//package main
//import (
//    "os"
//    "os/exec"
//)
//func main() {
//    c1 := exec.Command("ls")
//    c2 := exec.Command("wc", "-l")
//    c2.Stdin, _ = c1.StdoutPipe()
//    c2.Stdout = os.Stdout
//    _ = c2.Start()
//    _ = c1.Run()
//    _ = c2.Wait()
//}
