package models

import (
	"time"
)

//type Tag string

type Article struct {
	// 文章的标题
	Title string
	// 创建时间
	CreatedAt time.Time `toml:"created_at"`
	// 所属分类的名称
	Category string
	// 文章主题内容， markdown
	Body string
	// 文章文件路径
	Path string
	//文章标签
	Tags []Tag

	Description string
}

type ArticleInfo struct {
	Title string
	Category string	// 所属分类的名称
	CreatedAt time.Time
	Tags []Tag
	Description string
}

type ArticlesPagination struct {
	Articles []ArticleInfo
	Total int
	CurrentPage int
	PageNum []int
}

type Articles []ArticleInfo


func (a Articles) Len() int {
	return len(a)
}

func (a Articles) Less(i, j int) bool {
	return a[i].CreatedAt.After(a[j].CreatedAt)
}

func (a Articles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

