package log

type Log interface {
	Fatal(format string, v ...interface{})
	Error(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
}
