package service

import (
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/models"
	"math"
	"strings"
)

func GetArticleList(page int, dir string, search string) (models.MarkdownPagination, error) {

	allArticle, err := models.GetMarkdownListByCache(dir)

	if err != nil {
		return models.MarkdownPagination{}, err
	}
	if "" == search{
		return getPaginationData(allArticle, page)
	}

	var newArticleList models.MarkdownList

	for _, article := range allArticle {
		if strings.Index(article.Title,search) != -1{
			newArticleList = append(newArticleList,article)
		}
	}
	return  getPaginationData(newArticleList, page)
}

func getPaginationData(allArticle models.MarkdownList, page int) (models.MarkdownPagination, error) {

	var paginationData models.MarkdownPagination

	articleLen := len(allArticle)
	pageSize := config.Cfg.PageSize
	totalPage := int(math.Floor(float64(articleLen / pageSize)))

	if (articleLen % pageSize) != 0 {
		totalPage++
	}

	paginationData.Total = articleLen
	paginationData.CurrentPage = page
	paginationData.PageNumber = helper.BuildArrByInt(totalPage)

	if page < 1 || pageSize*(page-1) > articleLen { //超出页码

		paginationData.CurrentPage = 1

		if pageSize <= articleLen {
			paginationData.Markdowns = allArticle[0:pageSize]
		} else {
			paginationData.Markdowns = allArticle[0:articleLen]
		}
		return paginationData, nil
	}

	startNum := (page - 1) * pageSize
	endNum := startNum + pageSize

	if endNum > articleLen {
		paginationData.Markdowns = allArticle[startNum:articleLen]
	} else {
		paginationData.Markdowns = allArticle[startNum:endNum]
	}

	return paginationData, nil
}
