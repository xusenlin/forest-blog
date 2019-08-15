package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
	"time"
)

func GetArticleDetails(path string) (Article , error) {

	fullPath := config.Cfg.DocumentPath + "/" + path
	categoryName := strings.Split(path,"/")[1]
	markdownFile, fileErr := os.Stat(fullPath)

	var emptyArticle Article

	if fileErr != nil {
		return emptyArticle,fileErr
	}
	if markdownFile.IsDir() {
		return emptyArticle,errors.New("this path is Dir")
	}
	markdown, mdErr := GetMarkdownByPath(path)

	if mdErr != nil {
		return emptyArticle,mdErr
	}

	markdown = bytes.TrimSpace(markdown)

	markdownArticle := Article{
		markdownFile.Name(),
		markdownFile.ModTime(),
		categoryName,
		string(markdown),
		fullPath,
		nil,
		"",
	}

	if ! bytes.HasPrefix(markdown,[]byte("```json")) {
		return markdownArticle,nil
	}

	markdown = bytes.Replace(markdown,[]byte("```json"),[]byte(""),1)

	article := bytes.SplitN(markdown,[]byte("```"),2)

	var info struct{
		Date string `json:"date"`
		Description string `json:"description"`
		Tags []Tag `json:"tags"`
	}

	if json.Unmarshal(bytes.TrimSpace(article[0]),&info) != nil{
		return markdownArticle,nil
	}

	loc, _ := time.LoadLocation("Local")
	timeVal,_ := time.ParseInLocation("2006-01-02 15:04",info.Date,loc)

	return Article{
		markdownFile.Name(),
		timeVal,
		categoryName,
		string(article[1]),
		fullPath,
		info.Tags,
		info.Description,
	},nil

}

func GetArticles(page int , categoryName string) ArticlesPagination {

	var allArticle []ArticleInfo
	if len(categoryName) != 0 {
		allArticle = getArticleByCategoryName(categoryName)
	}else {
		allArticle = getAllArticle()
	}

	articleLen := len(allArticle)
	pageSize := config.Cfg.PageSize
	totalPage := int(math.Floor(float64(articleLen / pageSize)))

	if (articleLen % pageSize) != 0 {
		totalPage ++
	}

	pageNum := helper.BuildArrByInt(totalPage)

	if page < 1 || pageSize*(page-1) > articleLen { //超出页码

		if pageSize <= articleLen {
			article := allArticle[0:pageSize]
			return ArticlesPagination{article, articleLen, 1, pageNum}
		} else {
			article := allArticle[0:articleLen]
			return ArticlesPagination{article, articleLen, 1, pageNum}
		}
	}

	startNum := (page - 1) * pageSize
	endNum := startNum + pageSize

	if endNum > articleLen {
		article := allArticle[startNum:articleLen]
		return ArticlesPagination{article, articleLen, page, pageNum}
	} else {
		article := allArticle[startNum:endNum]
		return ArticlesPagination{article, articleLen, page, pageNum}
	}

}

func getAllArticle() []ArticleInfo {
	var allArticle Articles

	CategoriesDirs, _ := ioutil.ReadDir(config.Cfg.DocumentPath + "/content")
	for _, CategoriesDir := range CategoriesDirs {

		if CategoriesDir.IsDir() {

			CategoriesMdFile, _ := ioutil.ReadDir(config.Cfg.DocumentPath + "/content/" + CategoriesDir.Name())

			for _, markdownFile := range CategoriesMdFile {
				allArticle = append(allArticle,
					ArticleInfo{markdownFile.Name(), CategoriesDir.Name(), markdownFile.ModTime(),nil,""})
			}

		}
	}
	sort.Sort(allArticle)
	return allArticle
}

func getArticleByCategoryName(categoryName string) []ArticleInfo {

	var allArticle Articles

	CategoriesMdFile, _ := ioutil.ReadDir(config.Cfg.DocumentPath + "/content/" + categoryName)

	for _, markdownFile := range CategoriesMdFile {

		allArticle = append(allArticle,
			ArticleInfo{markdownFile.Name(), categoryName, markdownFile.ModTime(),nil,""})

	}
	sort.Sort(allArticle)

	return allArticle
}