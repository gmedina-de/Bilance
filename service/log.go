package service

type Log interface {
	Debug(message string, parameters ...interface{})
	Info(message string, parameters ...interface{})
	Warning(message string, parameters ...interface{})
	Error(message string, parameters ...interface{})
}
