package logger

import (
	"log"
	"os"
)

var (
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func Init() {
	infoHandle := os.Stdout
	errorHandle := os.Stderr

	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(infoHandle, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
