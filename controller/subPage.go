package controller

import (
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/models"
	"github.com/xusenlin/go_blog/service"
	"net/http"
	"strconv"
	"strings"
)

func Article(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	path := r.Form.Get("path")

	template, err := helper.HtmlTemplate("article")
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	article, err := models.GetMarkdownDetails(path)
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "文章详情",
		"Data":   article,
		"Config": config.Cfg,
	})

	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}
}

func CategoryArticle(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	template, err := helper.HtmlTemplate("category")

	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	categoryName := r.Form.Get("name")
	page, err := strconv.Atoi(r.Form.Get("page"))
	if err != nil {
		page = 1
	}
	content,err := service.GetArticleList(page, categoryName,"")
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  strings.Replace(categoryName,"/","",1),
		"Data":   content,
		"Config": config.Cfg,
	})
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}
}

