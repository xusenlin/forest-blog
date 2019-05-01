package models

import (
	"github.com/xusenlin/go_blog/config"
	"io/ioutil"
	"time"
)



type Category struct {
	Title string
	Number int
	CreatedAt time.Time
	Article []ArticleInfo
}

type ArticleInfo struct {
	Title string
	CreatedAt time.Time
}

func GetCategoriesInfo() []Category {

	var Categories  []Category


	CategoriesDirs ,_ := ioutil.ReadDir(config.Cfg.DocumentPath + "/content")

	for _, CategoriesDir := range CategoriesDirs {
		if CategoriesDir.IsDir() {

			CategoriesMdFile ,_ := ioutil.ReadDir(config.Cfg.DocumentPath + "/content/" + CategoriesDir.Name())
			var mdArticle  []ArticleInfo
			for num, markdownFile := range CategoriesMdFile {
				if num >= config.Cfg.MaxNumberArticleOfCategory{
					goto Loop
				}
				mdArticle = append(mdArticle,ArticleInfo{markdownFile.Name(),markdownFile.ModTime()})
			}
			Loop :
			Categories = append(Categories, Category{CategoriesDir.Name(),len(CategoriesMdFile),CategoriesDir.ModTime(),mdArticle})
		}
	}

	return Categories
}
