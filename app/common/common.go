package common

import (
	"genuine/app/common/controllers"
	"genuine/app/common/filters"
	_ "genuine/app/common/localizations"
	"genuine/app/common/navigation"
	"genuine/app/common/repositories"
	_ "genuine/app/common/template"
	"genuine/core"
)

func init() {
	core.Provide(controllers.Search)
	core.Provide(repositories.Users)
	core.Provide(filters.Basic)
	core.Provide(navigation.Standard)

}
