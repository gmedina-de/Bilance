package log

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"log"
	"os"
)

var logger = logs.GetBeeLogger()

type console struct {
	beeLogger *logs.BeeLogger
}

func Console() Log {
	// adapt orm logger
	l := log.New(os.Stdout, "", 0)
	l.SetFlags(0)
	l.SetOutput(new(ormLogAdapter))
	orm.DebugLog.Logger = l
	return logger
}

type ormLogAdapter struct {
}

func (writer *ormLogAdapter) Write(bytes []byte) (int, error) {
	return logger.Write(bytes)
}
