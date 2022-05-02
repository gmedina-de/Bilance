package log

type color string

const (
	Reset    color = "\x1b[0m"
	FgBlack  color = "\x1b[30m"
	FgRed    color = "\x1b[31m"
	FgGreen  color = "\x1b[32m"
	FgYellow color = "\x1b[33m"
	FgBlue   color = "\x1b[34m"
	FgPurple color = "\x1b[35m"
	FgCyan   color = "\x1b[36m"
	FgWhite  color = "\x1b[37m"
	BgBlack  color = "\x1b[40m"
	BgRed    color = "\x1b[41m"
	BgGreen  color = "\x1b[42m"
	BgYellow color = "\x1b[43m"
	BgBlue   color = "\x1b[44m"
	BgPurple color = "\x1b[45m"
	BgCyan   color = "\x1b[46m"
	BgWhite  color = "\x1b[47m"
)
