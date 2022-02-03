package template

type Pagination struct {
	Pages  int64
	Active int64
}

func paginate(count int64) []int64 {
	var i int64
	var items []int64
	for i = 1; i <= count; i++ {
		items = append(items, i)
	}
	return items
}
