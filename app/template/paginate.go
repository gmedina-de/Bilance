package template

func paginate(pages int64, page int64, offset int64) []int64 {
	var i int64
	var items []int64
	for i = page - offset; i <= page+offset; i++ {
		if i <= pages && i > 0 {
			items = append(items, i)
		}
	}
	return items
}
