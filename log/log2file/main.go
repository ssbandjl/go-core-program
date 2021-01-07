package main

import (
	"log"
	"os"
)

//自定义日志记录器,将日志记录到文件
type MyLogger struct {
	LogFile *os.File
	Log     *log.Logger
}

func (i *MyLogger) GetLogFile() (err error, file *os.File) {
	f, err := os.OpenFile("main.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	i.LogFile = f
	return nil, f
}

func (i *MyLogger) GetLogger() (err error, logger *log.Logger) {
	logger = log.New(i.LogFile, "", log.LstdFlags)
	i.Log = logger
	return nil, logger
}

func main() {
	//f, err := os.OpenFile("main.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//if err != nil {
	//	log.Println(err)
	//}
	//defer f.Close()
	//
	//logger := log.New(f, "prefix", log.LstdFlags)
	//logger.Println("text to append")
	//logger.Println("more text to append")
	//logger.Printf("日志到文件")
	myLogger := MyLogger{}
	myLogger.GetLogFile()
	myLogger.GetLogger()
	defer myLogger.LogFile.Close()
	myLogger.Log.Printf("日志记录")
}
