package template

import (
	"github.com/beego/beego/v2/server/web"
	"homecloud/core/localization"
	"strconv"
	"strings"
)

func init() {
	web.AddFuncMap("translate", localization.Translate)
	web.AddFuncMap("paginate", paginate)
	web.AddFuncMap("inputs", inputs)
	web.AddFuncMap("td", td)
	web.AddFuncMap("th", th)
	web.AddFuncMap("sum", func(a int, b int) int { return a + b })
	web.AddFuncMap("contains", func(a string, b int64) bool { return strings.Contains(a, strconv.FormatInt(b, 10)) })
}
