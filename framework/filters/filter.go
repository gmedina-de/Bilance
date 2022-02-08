package filters

import "github.com/beego/beego/v2/server/web"

type Filter interface {
	Pattern() string
	Pos() int
	Func() web.FilterFunc
	Opts() []web.FilterOpt
}
