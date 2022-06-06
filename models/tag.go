package models

import (
	"sort"
)

type Tag struct {
	Name     string
	Quantity int
	Articles Articles
}
type Tags []Tag

func GroupByTag(articles *Articles, articleQuantity int) Tags {
	var tags Tags
	tagMap := make(map[string]Articles)

	for _, article := range *articles {
		for _, tag := range article.Tags {
			_, existedCategory := tagMap[tag]
			if existedCategory {
				tagMap[tag] = append(tagMap[tag], article)
			} else {
				tagMap[tag] = Articles{article}
			}
		}
	}
	for categoryName, articleItem := range tagMap {
		articleLen := len(articleItem)

		var articleList Articles
		if articleQuantity <= 0 {
			articleList = articleItem
		} else {
			if articleQuantity > articleLen {
				articleList = articleItem[0:articleLen]
			} else {
				articleList = articleItem[0:articleQuantity]
			}
		}
		tags = append(tags, Tag{
			Name:     categoryName,
			Quantity: articleLen,
			Articles: articleList,
		})
	}
	sort.Sort(tags)
	return tags
}

func (c Tags) Len() int { return len(c) }

func (c Tags) Less(i, j int) bool { return c[i].Quantity > c[j].Quantity }

func (c Tags) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
