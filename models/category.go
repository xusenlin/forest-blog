package models

import (
	"ForestBlog/config"
	"sort"
	"strings"
)

type Category struct {
	Name     string
	Quantity int
	Articles Articles
}
type Categories []Category

func GetCategoryName(path string) string {
	var categoryName string
	newPath := strings.Replace(path, config.Cfg.DocumentContentDir+"/", "", 1)

	if strings.Index(newPath, "/") == -1 { //文件在根目录下(content/)没有分类名称
		categoryName = "未分类"
	} else {
		categoryName = strings.Split(newPath, "/")[0]
	}
	return categoryName
}

func GroupByCategory(articles *Articles, articleQuantity int) Categories {

	var categories Categories
	categoryMap := make(map[string]Articles)

	for _, article := range *articles {

		_, existedCategory := categoryMap[article.Category]
		if existedCategory {
			categoryMap[article.Category] = append(categoryMap[article.Category], article)
		} else {
			categoryMap[article.Category] = Articles{article}
		}
	}
	for categoryName, articles := range categoryMap {
		articleLen := len(articles)

		var articleList Articles
		if articleQuantity <= 0 {
			articleList = articles
		} else {
			if articleQuantity > articleLen {
				articleList = articles[0:articleLen]
			} else {
				articleList = articles[0:articleQuantity]
			}
		}
		categories = append(categories, Category{
			Name:     categoryName,
			Quantity: articleLen,
			Articles: articleList,
		})
	}
	sort.Sort(categories)
	return categories
}

func (c Categories) Len() int { return len(c) }

func (c Categories) Less(i, j int) bool { return c[i].Quantity > c[j].Quantity }

func (c Categories) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
