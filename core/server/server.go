package server

var (
	Port = 8080
)

type Server interface {
	Start()
}
