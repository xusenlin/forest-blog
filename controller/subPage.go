package controller

import (
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/models"
	"net/http"
	"strconv"
)

func Article(w http.ResponseWriter, r *http.Request)  {

	_ = r.ParseForm()

	path := r.Form.Get("path")

	article,readErr := models.GetArticleDetails("content/" + path)
	if readErr != nil {
		_,_ = w.Write(helper.ErrorHtml(readErr.Error()))
		return
	}

	template, templateErr := helper.HtmlTemplate("article")
	if templateErr != nil {
		_,_ = w.Write(helper.ErrorHtml(templateErr.Error()))
		return
	}

	err := template.Execute(w, map[string]interface{}{
		"Title":"文章详情",
		"Data": article,
		"Config":config.Cfg,
	})
	if err != nil {
		_,_ = w.Write(helper.ErrorHtml(err.Error()))
		return
	}
}

func CategoryArticle(w http.ResponseWriter, r *http.Request)  {
	_ = r.ParseForm()

	categoryName := r.Form.Get("name")
	page,pageErr := strconv.Atoi(r.Form.Get("page"))
	if pageErr != nil{
		page = 1
	}

	template, templateErr := helper.HtmlTemplate("category")
	if templateErr != nil {
		_,_ = w.Write(helper.ErrorHtml(templateErr.Error()))
		return
	}

	err := template.Execute(w, map[string]interface{}{
		"Title":categoryName,
		"Data": models.GetArticles(page,categoryName),
		"Config":config.Cfg,
	})
	if err != nil {
		_,_ = w.Write(helper.ErrorHtml(err.Error()))
		return
	}
}