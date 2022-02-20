package log

import (
	"fmt"
	"genuine/core/config"
	"log"
	"os"
	"strconv"
	"strings"
)

type standard struct {
	level Level
}

func Console() Log {
	return &standard{level: config.LogLevel}
}

func (s *standard) Fatal(tag Tag, format string, v ...interface{}) {
	s.log(tag, Fatal, format, v)
	os.Exit(-1)
}

func (s *standard) Error(tag Tag, format string, v ...interface{}) {
	s.log(tag, Error, format, v)
}

func (s *standard) Warning(tag Tag, format string, v ...interface{}) {
	s.log(tag, Warning, format, v)
}

func (s *standard) Info(tag Tag, format string, v ...interface{}) {
	s.log(tag, Info, format, v)
}

func (s *standard) Debug(tag Tag, format string, v ...interface{}) {
	s.log(tag, Debug, format, v)
}

const trim = 6

func (s *standard) log(tag Tag, level Level, format string, v []interface{}) {
	if len(tag) > trim {
		tag = tag[:trim]
	}
	if level <= s.level {
		log.Printf("%s %"+strconv.Itoa(trim)+"s %s%s %s%s %s\n",
			level.toBgColor(),
			strings.ToUpper(string(tag)),
			Reset,
			level.toFgColor(),
			level.String(),
			Reset,
			fmt.Sprintf(format, v...),
		)
	}
}
