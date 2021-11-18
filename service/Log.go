package service

import (
	"fmt"
	"time"
)

type Log interface {
	Debug(message string, parameters ...interface{})
	Info(message string, parameters ...interface{})
	Warning(message string, parameters ...interface{})
	Error(message string, parameters ...interface{})
}

const (
	LogLevelSetting Setting = iota
)

type AnsiColor string

const (
	Reset    AnsiColor = "\x1b[0m"
	FgBlack  AnsiColor = "\x1b[30m"
	FgRed    AnsiColor = "\x1b[31m"
	FgGreen  AnsiColor = "\x1b[32m"
	FgYellow AnsiColor = "\x1b[33m"
	FgBlue   AnsiColor = "\x1b[34m"
	FgPurple AnsiColor = "\x1b[35m"
	FgCyan   AnsiColor = "\x1b[36m"
	FgWhite  AnsiColor = "\x1b[37m"
	BgBlack  AnsiColor = "\x1b[40m"
	BgRed    AnsiColor = "\x1b[41m"
	BgGreen  AnsiColor = "\x1b[42m"
	BgYellow AnsiColor = "\x1b[43m"
	BgBlue   AnsiColor = "\x1b[44m"
	BgPurple AnsiColor = "\x1b[45m"
	BgCyan   AnsiColor = "\x1b[46m"
	BgWhite  AnsiColor = "\x1b[47m"
)

type Level string

const (
	Debug   Level = "DBUG"
	Info    Level = "INFO"
	Warning Level = "WARN"
	Error   Level = "ERRO"
)

func (l Level) toFgColor() AnsiColor {
	switch l {
	case Debug:
		return FgGreen
	case Info:
		return FgBlue
	case Warning:
		return FgYellow
	case Error:
		return FgRed
	default:
		return FgGreen
	}
}

func (l Level) toBgColor() AnsiColor {
	switch l {
	case Debug:
		return BgGreen
	case Info:
		return BgBlue
	case Warning:
		return BgYellow
	case Error:
		return BgRed
	default:
		return BgGreen
	}
}

type consoleLog struct {
	configuration Configuration
}

func ConsoleLog(configuration Configuration) Log {
	return &consoleLog{configuration: configuration}
	// TODO use built in go logger?
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
	fmt.Printf("%s %s %s%s %s%s %s\n",
		time.Now().Format("2006-01-02 15:04:05"),
		level.toBgColor(),
		Reset,
		level.toFgColor(),
		level,
		Reset,
		fmt.Sprintf(message, parameters...),
	)
}
