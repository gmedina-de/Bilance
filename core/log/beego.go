package log

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"log"
	"os"
	"time"
)

func Beego() Log {
	logger := logs.GetBeeLogger()

	// adapt orm logger
	l := log.New(os.Stdout, "", 0)
	l.SetFlags(0)
	l.SetOutput(new(ormLogAdapter))

	orm.DebugLog.Logger = l
	return logger
}

type ormLogAdapter struct {
}

const d = " \033[1;44m[D]\033[0m "
const l = "2006/01/02 15:04:05.000"

func (writer *ormLogAdapter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().UTC().Format(l) + d + string(bytes))
}
