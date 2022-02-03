package log

import (
	"fmt"
	"log"
)

type console struct {
}

func Console() Log {
	return &console{}
}

func (this *console) Debug(message string, parameters ...interface{}) {
	this.log(Debug, message, parameters...)
}

func (this *console) Info(message string, parameters ...interface{}) {
	this.log(Info, message, parameters...)
}

func (this *console) Warning(message string, parameters ...interface{}) {
	this.log(Warning, message, parameters...)
}

func (this *console) Error(message string, parameters ...interface{}) {
	this.log(Error, message, parameters...)
}

func (this *console) log(level LogLevel, message string, parameters ...interface{}) {
	log.Printf("%s %s%s %s%s %s\n",
		level.toBgColor(),
		Reset,
		level.toFgColor(),
		level,
		Reset,
		fmt.Sprintf(message, parameters...),
	)
}
