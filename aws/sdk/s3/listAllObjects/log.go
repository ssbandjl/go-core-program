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
