package logger

import (
	"log"
	"os"
)

var (
	info *log.Logger
	warn *log.Logger
	err  *log.Logger
)

func Init() {
	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warn = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	err = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(msg string) {
	info.Println(msg)
}

func Warning(msg string) {
	warn.Println(msg)
}

func Error(msg string) {
	err.Println(msg)
}
