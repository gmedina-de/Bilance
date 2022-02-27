package common

import (
	"genuine/app/common/authenticator"
	_ "genuine/app/common/localizations"
	"genuine/app/common/navigation"
	"genuine/app/common/repositories"
	_ "genuine/app/common/template"
	"genuine/core"
)

func init() {
	core.Provide(repositories.Users)
	core.Provide(authenticator.Basic)
	core.Provide(navigation.Standard)

}
