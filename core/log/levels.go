package log

type Level int

const (
	Fatal   Level = iota
	Error   Level = iota
	Warning Level = iota
	Info    Level = iota
	Debug   Level = iota
)

func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFOM"
	case Warning:
		return "WARNG"
	case Error:
		return "ERROR"
	default:
		return "FATAL"
	}
}

func (l Level) toFgColor() color {
	switch l {
	case Debug:
		return FgGreen
	case Info:
		return FgBlue
	case Warning:
		return FgYellow
	default:
		return FgRed
	}
}

func (l Level) toBgColor() color {
	switch l {
	case Debug:
		return BgGreen
	case Info:
		return BgBlue
	case Warning:
		return BgYellow
	default:
		return BgRed
	}
}
