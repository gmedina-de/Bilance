package log

type Log interface {
	Fatal(tag Tag, format string, v ...interface{})
	Error(tag Tag, format string, v ...interface{})
	Warning(tag Tag, format string, v ...interface{})
	Info(tag Tag, format string, v ...interface{})
	Debug(tag Tag, format string, v ...interface{})
}

type Tag = string
