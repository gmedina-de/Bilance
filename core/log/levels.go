package log

type Level = int

const (
	Critical Level = iota
	Error    Level = iota
	Warning  Level = iota
	Info     Level = iota
	Debug    Level = iota
)
