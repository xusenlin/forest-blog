package models

import (
	"github.com/xusenlin/go_blog/config"
	"io/ioutil"
	"sort"
	"time"
)

type Article struct {
	// 文章的标题
	Title string
	// 创建时间
	CreatedAt time.Time `toml:"created_at"`
	// 所属分类的名称
	Category string
	// 文章主题内容， markdown
	Body string
	// 文章在服务器上的文件路由
	Path string
}

type ArticleInfo struct {
	Title string
	Category string	// 所属分类的名称
	CreatedAt time.Time
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

func GetArticleByPage(page int) []ArticleInfo {

	article := getAllArticle()
	articleLen := len(article)

	if page < 1 || config.Cfg.PageSize * (page-1) > articleLen{//超出页码

		if config.Cfg.PageSize <= articleLen{
			return article[0 : config.Cfg.PageSize]
		}else {
			return article[0 : articleLen]
		}
	}

	startNum := (page-1) * config.Cfg.PageSize
	endNum := startNum + config.Cfg.PageSize

	if endNum > articleLen {
		return article[startNum : articleLen]
	}else {
		return  article[startNum : endNum]
	}

}


func getAllArticle() []ArticleInfo {
	var allArticle Articles

	CategoriesDirs ,_ := ioutil.ReadDir(config.Cfg.DocumentPath + "/content")
	for _, CategoriesDir := range CategoriesDirs {

		if CategoriesDir.IsDir() {

			CategoriesMdFile ,_ := ioutil.ReadDir(config.Cfg.DocumentPath + "/content/" + CategoriesDir.Name())

			for _, markdownFile := range CategoriesMdFile {
				allArticle = append(allArticle,ArticleInfo{markdownFile.Name(),CategoriesDir.Name(),markdownFile.ModTime()})
			}

		}
	}
	sort.Sort(allArticle)
	return allArticle
}

