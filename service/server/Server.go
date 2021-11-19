package server

import (
	"Bilance/service/configuration"
)

type Server interface {
	Start()
	Stop()
}

const (
	Port configuration.Setting = iota
)
