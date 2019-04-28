package utils

import (
	"net/http"
	"strconv"
)

type PageResponse struct {
	Data          interface{} `json:"data"`
	Limit         int         `json:"limit"`
	Elements      int         `json:"elements"`
	TotalElements int         `json:"total_elements"`
	CurrentPage   int         `json:"current_page"`
	Pages         int         `json:"pages"`
}

func PageRequest(r *http.Request, limit int) (int, int, int) {
	keys := r.URL.Query()
	if keys.Get("page") == "" {
		return 1, 0, limit
	}
	page, _ := strconv.Atoi(keys.Get("page"))
	if page < 1 {
		return 1, 0, limit
	}
	init := (limit * page) - limit
	return page, init, limit
}

func Pagination(data interface{}, elements, limit, total, page int) interface{} {
	pages := PageCount(total, limit)
	if pages == 0 {
		return data
	}
	return PageResponse{data, limit, elements, total, page, pages}
}

func PageCount(total, limit int) int {
	pages := (total / limit)
	if (total % limit) != 0 {
		pages++
	}
	return pages
}