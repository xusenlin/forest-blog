package models

import (
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"io/ioutil"
	"math"
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

func GetArticleByPage(page int) ArticlesPagination {

	article := getAllArticle()
	articleLen := len(article)
	pageSize := config.Cfg.PageSize
	totalPage := int(math.Floor(float64(articleLen / pageSize)))

	if (articleLen % pageSize) != 0 {
		totalPage ++
	}

	pageNum := helper.BuildArrByInt(totalPage)

	if page < 1 || pageSize * (page-1) > articleLen{//超出页码

		if pageSize <= articleLen{
			article := article[0 : pageSize]
			return ArticlesPagination{article,articleLen,1,pageNum}
		}else {
			article := article[0 : articleLen]
			return ArticlesPagination{article,articleLen,1,pageNum}
		}
	}

	startNum := (page-1) * pageSize
	endNum := startNum + pageSize

	if endNum > articleLen {
		article := article[startNum : articleLen]
		return ArticlesPagination{article,articleLen,page,pageNum}
	}else {
		article := article[startNum : endNum]
		return  ArticlesPagination{article,articleLen,page,pageNum}
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

