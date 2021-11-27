package model

import (
	"strconv"
	"strings"
)

type Category struct {
	Id        int64
	Name      string
	Color     string
	ProjectId int64
}

func ExtractCategoryIds(categories []Category) string {
	var result []string
	for _, category := range categories {
		result = append(result, strconv.FormatInt(category.Id, 10))
	}
	return strings.Join(result, ",")
}
