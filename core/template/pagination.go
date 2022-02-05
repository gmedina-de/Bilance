package template

import (
	"math"
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

func HandlePagination(request *http.Request, count int) (int, int, *Pagination) {
	var limit = 10
	var page = 1
	if request.URL.Query().Get("page") != "" {
		page, _ = strconv.Atoi(request.URL.Query().Get("page"))
	}
	var pages = int(count-1) / limit
	if page == 0 {
		limit = math.MaxInt
	}
	var offset = limit * (page - 1)
	pages++
	pagination := &Pagination{pages, page}
	return limit, offset, pagination
}
