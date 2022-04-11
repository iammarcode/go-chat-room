package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	File *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	Logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	Logger.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	Logger.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	Logger.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	Logger.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	Logger.Fatalln(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	Logger.SetPrefix(logPrefix)
}