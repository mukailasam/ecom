package logger

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
	WarningLog *log.Logger
}

func LoggerFile() (*os.File, error) {
	f, err := os.OpenFile("ecom.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func NewLogger() *Logger {
	logFile, err := LoggerFile()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	infoLogger := log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	errorLogger := log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
	warningLogger := log.New(logFile, "WARNING: ", log.Ldate|log.Ltime|log.Llongfile)

	l := &Logger{
		InfoLog:    infoLogger,
		ErrorLog:   errorLogger,
		WarningLog: warningLogger,
	}

	return l
}
