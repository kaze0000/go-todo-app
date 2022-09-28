package utils

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)//ファイルの読み書き、作成、追記を指定
	if err != nil{
		log.Fatalln(err)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)//logの書き込み先を標準出力とlogfileに指定
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	log.SetOutput(multiLogFile)
}
