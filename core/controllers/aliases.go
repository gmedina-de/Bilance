package controllers

import (
	"genuine/core/models"
	"net/http"
	"strconv"
)

type Request struct {
	*http.Request
	http.ResponseWriter
}

func (r *Request) GetID() uint {
	parseUint, _ := strconv.ParseUint(r.URL.Query().Get(models.ID), 10, 64)
	return uint(parseUint)
}

type Response = map[string]any

type Handler func(r Request) Response

type Routes map[string]any

type Template string
