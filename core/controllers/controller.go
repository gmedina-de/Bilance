package controllers

type Controller interface {
	Routes() map[string]Handler
}
