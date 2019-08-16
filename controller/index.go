package controller

import (
	"fmt"
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
	}

	page, err := strconv.Atoi(r.Form.Get("page"))
	if err != nil {
		page = 1
	}

	template, err := helper.HtmlTemplate("index")
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
	}

	markdownPagination, err := service.GetArticleList(page, "/")
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "首页",
		"Data":   markdownPagination,
		"Config": config.Cfg,
	})

	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
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
	fmt.Println(categories)
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}
}

func Works(w http.ResponseWriter, r *http.Request) {

	markdown, err := models.GetMarkdownDetails("/Works.md")
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
	}

	template, err := helper.HtmlTemplate("works")
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "作品",
		"Body":   markdown.Body,
		"Config": config.Cfg,
	})
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
	}
}

//
//func About(w http.ResponseWriter, r *http.Request)  {
//
//	markdown,mdErr := models.GetMarkdownByPath("About.md")
//	if mdErr != nil {
//		_,_ = w.Write(helper.ErrorHtml(mdErr.Error()))
//		return
//	}
//
//	template, templateErr := helper.HtmlTemplate("about")
//	if templateErr != nil {
//		_,_ = w.Write(helper.ErrorHtml(templateErr.Error()))
//		return
//	}
//
//	err := template.Execute(w, map[string]interface{}{
//		"Title": "关于",
//		"Data": string(markdown),
//		"Config":config.Cfg,
//	})
//	if err != nil {
//		_,_ = w.Write(helper.ErrorHtml(err.Error()))
//		return
//	}
//
//}
