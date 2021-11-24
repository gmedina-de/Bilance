package service

type Server interface {
	Start()
	Stop()
}

const (
	Port Setting = iota
)
