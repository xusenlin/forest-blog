package models

import "time"

type Tag string

type MarkdownBaseInfo struct {
	Title string `json:"title"`
	Date string `json:"date"`
	Description string `json:"description"`
	Tags []Tag `json:"tags"`
	Author string `json:"author"`
}

type Markdown struct {
	Title string
	Author string
	Category string
	CreatedAt time.Time
	Tags []Tag
	Description string
}
type MarkdownList []Markdown

type MarkdownPagination struct {
	Markdowns MarkdownList
	Total int
	CurrentPage int
	PageNum []int
}

type MarkdownDetails struct {
	Title string
	CreatedAt time.Time
	Category string
	Body string
	Path string
	Tags []Tag
	Description string
}