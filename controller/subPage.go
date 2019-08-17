package controller

import (
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/models"
	"net/http"
)

func Article(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
	}

	path := r.Form.Get("path")

	template, err := helper.HtmlTemplate("article")
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
	}

	article, err := models.GetMarkdownDetails(path)
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "文章详情",
		"Data":   article,
		"Config": config.Cfg,
	})

	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
	}
}
//
//func CategoryArticle(w http.ResponseWriter, r *http.Request) {
//	_ = r.ParseForm()
//
//	categoryName := r.Form.Get("name")
//	page, pageErr := strconv.Atoi(r.Form.Get("page"))
//	if pageErr != nil {
//		page = 1
//	}
//
//	template, templateErr := helper.HtmlTemplate("category")
//	if templateErr != nil {
//		_, _ = w.Write(helper.ErrorHtml(templateErr.Error()))
//		return
//	}
//
//	err := template.Execute(w, map[string]interface{}{
//		"Title":  categoryName,
//		"Data":   models.GetArticles(page, categoryName),
//		"Config": config.Cfg,
//	})
//	if err != nil {
//		_, _ = w.Write(helper.ErrorHtml(err.Error()))
//		return
//	}
//}
