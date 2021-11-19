package log

import (
	"fmt"
	"log"
	"time"
)

type consoleLog struct {
}

func ConsoleLog() Log {
	return &consoleLog{}
}

func (this *consoleLog) Debug(message string, parameters ...interface{}) {
	this.log(Debug, message, parameters...)
}

func (this *consoleLog) Info(message string, parameters ...interface{}) {
	this.log(Info, message, parameters...)
}

func (this *consoleLog) Warning(message string, parameters ...interface{}) {
	this.log(Warning, message, parameters...)
}

func (this *consoleLog) Error(message string, parameters ...interface{}) {
	this.log(Error, message, parameters...)
}

func (this *consoleLog) log(level Level, message string, parameters ...interface{}) {
	log.Printf("%s %s %s%s %s%s %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		level.toBgColor(),
		Reset,
		level.toFgColor(),
		level,
		Reset,
		fmt.Sprintf(message, parameters...),
	)
}
