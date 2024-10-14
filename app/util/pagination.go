package util

import (
	"fmt"
	"net/url"
	"strconv"
)

const (
	MAX_PAGE_SIZE = 99999999
)

type Pagination struct {
	Page     int64 `json:"page"`
	PageSize int64 `json:"pageSize"`
}

func GetPageInQuery(query url.Values) (*Pagination, error) {
	paginagion := Pagination{}
	fmt.Println(query.Get("page"), query.Get("pageSize"))
	page, err := strconv.ParseInt(query.Get("page"), 10, 64)
	if err != nil || query.Get("page") == "" {
		page = 1
	}

	pageSize, err := strconv.ParseInt(query.Get("pageSize"), 10, 64)
	if err != nil || query.Get("pageSize") == "" {
		pageSize = MAX_PAGE_SIZE
	}
	paginagion.Page = page
	paginagion.PageSize = pageSize
	return &paginagion, nil
}
