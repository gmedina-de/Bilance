package template

import (
	"net/http"
	"strconv"
)

type Pagination struct {
	Pages  int
	Active int
}

func paginate(count int) []int {
	var i int
	var items []int
	for i = 1; i <= count; i++ {
		items = append(items, i)
	}
	return items
}

func HandlePagination(request *http.Request) (int, int, *Pagination) {
	var limit = 10
	var page = 1
	if request.URL.Query().Get("page") != "" {
		page, _ = strconv.Atoi(request.URL.Query().Get("page"))
	}
	var offset = limit * (page - 1)
	var pages = 10
	pages++
	pagination := &Pagination{pages, page}
	return limit, offset, pagination
}
