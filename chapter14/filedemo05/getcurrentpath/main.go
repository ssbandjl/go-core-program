//https://segmentfault.com/a/1190000013685370 Golang “相对”路径问题
//go run 是在临时文件夹下运行的
package main

import (
    "path/filepath"
    "os"
    "os/exec"
    "strings"
    "log"
)

func GetAppPath() string {
    file, _ := exec.LookPath(os.Args[0])
    path, _ := filepath.Abs(file)
    index := strings.LastIndex(path, string(os.PathSeparator))

    return path[:index]
}

func main(){
    log.Println(GetAppPath())
    //2020/03/28 20:40:26 C:\Users\XB\AppData\Local\Temp\go-build777338702\command-line-arguments\_obj\exe
}