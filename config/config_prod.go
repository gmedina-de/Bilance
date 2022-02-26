//go:build prod

package config

const (
	AppName        = "gCloud"
	ServerPort     = 80
	LogLevel       = 3 //log.Info
	ViewsDirectory = "views"
	ViewExtension  = ".gohtml"
)
