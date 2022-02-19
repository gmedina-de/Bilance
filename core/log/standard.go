package log

import (
	"github.com/beego/beego/v2/core/logs"
)

var logger = logs.GetBeeLogger()

type standard struct {
}

func Console() Log {
	return &standard{}
}

func (s *standard) Critical(format string, v ...interface{}) {
	s.log(format, v)
}

func (s *standard) Error(format string, v ...interface{}) {
	s.log(format, v)
}

func (s *standard) Warning(format string, v ...interface{}) {
	s.log(format, v)
}

func (s *standard) Info(format string, v ...interface{}) {
	s.log(format, v)
}

func (s *standard) Debug(format string, v ...interface{}) {
	s.log(format, v)
}

func (s *standard) log(format string, v []interface{}) {
	logger.Debug(format, v)
}
