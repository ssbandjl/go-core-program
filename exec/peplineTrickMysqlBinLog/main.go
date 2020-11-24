package main

//上面的解决方案是Go风格的解决方案，事实上你还可以用一个"Trick"花招/hack方式/来实现。

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("mysqlbinlog", "-vv", "mysql-bin.002400")
	//cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Printf("执行StdoutPipe出错:%s", err.Error())
	}

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		log.Printf("执行Start出错:%s", err.Error())
	}

	//time.Sleep(time.Second)
	// read command's stdout line by line
	in := bufio.NewScanner(stdout)

	//创建一个新文件，写入内容 5句 "hello, Gardon"
	//1 .打开文件 d:/abc.txt
	filePath := "abc.sql"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}
	//及时关闭file句柄
	defer file.Close()
	//准备写入5句 "hello, Gardon"
	//str := "hello,Gardon\r\n" // \r\n 表示换行,默认记事本认为\r是换行，其他编辑器可能认为\n才是换行
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)

	for in.Scan() {
		//len, err := writer.WriteString(in.Text())
		len, err := writer.WriteString(in.Text() + "\r\n")
		if err != nil {
			log.Printf("写入文件错误:%s,长度:%d", err.Error(), len)
		}
		//log.Printf(in.Text()) // write each line to your log, or anything you need
	}

	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer.Flush()

	if err := in.Err(); err != nil {
		log.Printf("error: %s", err)
	}

}
