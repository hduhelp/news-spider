package logger

import (
	"io"
	"log"
	"os"
)

const (
	ErrorLogPath = "errors.log"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	errFile, err := os.OpenFile(ErrorLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败！")
	}

	Info = log.New(os.Stdout, "[Info] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "[Warning] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "[Error] ", log.Ldate|log.Ltime|log.Lshortfile)
}
