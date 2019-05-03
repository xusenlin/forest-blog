package controller

import (

	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/models"
	"net/http"
	"strconv"
)

func Article(w http.ResponseWriter, r *http.Request)  {

	type ArticleInf struct {
		Title string
		Article models.Article
	}

	r.ParseForm()

	path := r.Form.Get("path")


	article := models.GetArticle("content/" + path)


	template, templateErr := helper.HtmlTemplate("article")
	if templateErr != nil {
		panic(templateErr)
	}

	err := template.Execute(w, map[string]ArticleInf{"Data": {"文章详情",article}})
	if err != nil {
		panic(err)
	}
}

func CategoryArticle(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	type CategoryInfo struct {
		Title string
		CategoryName string
		Pagination models.ArticlesPagination
	}

	categoryName := r.Form.Get("name")
	page,pageErr := strconv.Atoi(r.Form.Get("page"))
	if pageErr != nil{
		page = 1
	}
	categoryInfoData := CategoryInfo{"首页",categoryName,models.GetArticles(page,categoryName)}

	template, templateErr := helper.HtmlTemplate("category")
	if templateErr != nil {
		panic(templateErr)
	}

	err := template.Execute(w, map[string]CategoryInfo{"Data": categoryInfoData})
	if err != nil {
		panic(err)
	}
}