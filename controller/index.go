package controller

import (
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/models"
	"github.com/xusenlin/go_blog/service"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	page, err := strconv.Atoi(r.Form.Get("page"))
	if err != nil {
		page = 1
	}

	template, err := helper.HtmlTemplate("index")
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	searchKey := r.Form.Get("search")

	markdownPagination, err := service.GetArticleList(page, "/", searchKey)
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "首页",
		"Data":   markdownPagination,
		"Config": config.Cfg,
	})

	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

}

func Categories(w http.ResponseWriter, r *http.Request) {

	template, err := helper.HtmlTemplate("categories")

	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	categories, err := service.GetCategories()
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}
	err = template.Execute(w, map[string]interface{}{
		"Title":  "分类",
		"Data":   categories,
		"Config": config.Cfg,
	})

	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}
}

func Works(w http.ResponseWriter, r *http.Request) {

	template, err := helper.HtmlTemplate("works")
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	markdown, err := models.ReadMarkdownBody("/Works.md")
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "作品",
		"Data":   markdown,
		"Config": config.Cfg,
	})

	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

}

func About(w http.ResponseWriter, r *http.Request) {

	template, err := helper.HtmlTemplate("about")
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	markdown, err := models.ReadMarkdownBody("/About.md")

	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "关于",
		"Data":   markdown,
		"Config": config.Cfg,
	})
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}
}
