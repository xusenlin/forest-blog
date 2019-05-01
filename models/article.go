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

	// 作者姓名
	Author string

	// 创建时间
	CreatedAt time.Time `toml:"created_at"`

	// 最后更新时间
	UpdatedAt time.Time `toml:"updated_at"`

	// 标签
	Tags []string

	// 所属分类的名称
	Category string

	// 头部图片 URL 地址
	HeadImg string `toml:"head_img"`

	// 作者的个人主页
	HomePage string `toml:"home_page"`

	// 简短的描述
	Description string

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

func GetAllArticle() []ArticleInfo {
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
