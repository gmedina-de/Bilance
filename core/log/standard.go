package log

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var LogLevel = flag.Int("log", 4, "log level where 0 is critical and 4 is debug")

type standard struct {
	level Level
}

func Standard() Log {
	return &standard{level: Level(*LogLevel)}
}

func (s *standard) Critical(format string, v ...interface{}) {
	s.log(Critical, format, v)
	os.Exit(-1)
}

func (s *standard) Error(format string, v ...interface{}) {
	s.log(Error, format, v)
}

func (s *standard) Warning(format string, v ...interface{}) {
	s.log(Warning, format, v)
}

func (s *standard) Info(format string, v ...interface{}) {
	s.log(Info, format, v)
}

func (s *standard) Debug(format string, v ...interface{}) {
	s.log(Debug, format, v)
}

func (s *standard) log(level Level, format string, v []interface{}) {

	if level <= s.level {
		log.Printf("%s %s%s %s%s %s\n",
			level.toBgColor(),
			Reset,
			level.toFgColor(),
			level.String(),
			Reset,
			fmt.Sprintf(format, v...),
		)
	}
}
