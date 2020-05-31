package models

import (
	"math"
)

type PageResult struct {
	List      Articles `json:"list"`
	Total     int      `json:"total"`
	Page      int      `json:"page"`
	PageSize  int      `json:"pageSize"`
	TotalPage int
}

func Pagination(articles *Articles, page int, pageSize int) PageResult {

	articleLen := len(*articles)
	totalPage := int(math.Floor(float64(articleLen / pageSize)))

	if (articleLen % pageSize) != 0 {
		totalPage++
	}
	result := PageResult{
		Total:     articleLen,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: totalPage,
	}
	if page < 1 {
		result.Page = 1
	}
	if page > result.TotalPage {
		result.Page  = result.TotalPage
	}

	if articleLen <= result.PageSize {
		result.List = (*articles)[0:articleLen]
	} else {
		startNum := (result.Page - 1) * result.PageSize
		endNum := startNum + result.PageSize
		if endNum > articleLen {
			endNum = articleLen
		}
		result.List = (*articles)[startNum:endNum]
	}

	return result
}
