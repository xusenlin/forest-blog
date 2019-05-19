package controller

import (
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/models"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	page,pageErr := strconv.Atoi(r.Form.Get("page"))
	if pageErr != nil{
		page = 1
	}

	template, templateErr := helper.HtmlTemplate("index")
	if templateErr != nil {
		w.Write(helper.ErrorHtml(templateErr.Error()))
		return
	}

	err := template.Execute(w, map[string]interface{}{
		"Title":"首页",
		"Data": models.GetArticles(page,""),
		"Config":config.Cfg,
	})

	if err != nil {
		w.Write(helper.ErrorHtml(err.Error()))
		return
	}
}

func Categories(w http.ResponseWriter, r *http.Request)  {

	template, templateErr := helper.HtmlTemplate("categories")

	if templateErr != nil {
		w.Write(helper.ErrorHtml(templateErr.Error()))
		return
	}

	err := template.Execute(w, map[string]interface{}{
		"Title":"分类",
		"Data": models.GetCategoriesInfo(),
		"Config":config.Cfg,
	})
	if err != nil {
		w.Write(helper.ErrorHtml(err.Error()))
		return
	}
}

func Works(w http.ResponseWriter, r *http.Request)  {

	markdown,mdErr := models.GetMarkdownByPath("Works.md")
	if mdErr != nil {
		w.Write(helper.ErrorHtml(mdErr.Error()))
		return
	}

	template, templateErr := helper.HtmlTemplate("works")

	if templateErr != nil {
		w.Write(helper.ErrorHtml(templateErr.Error()))
		return
	}

	err := template.Execute(w, map[string]interface{}{
		"Title": "作品",
		"Body": string(markdown),
		"Config":config.Cfg,
	})
	if err != nil {
		w.Write(helper.ErrorHtml(err.Error()))
		return
	}
}

func About(w http.ResponseWriter, r *http.Request)  {

	markdown,mdErr := models.GetMarkdownByPath("About.md")
	if mdErr != nil {
		w.Write(helper.ErrorHtml(mdErr.Error()))
		return
	}

	template, templateErr := helper.HtmlTemplate("about")
	if templateErr != nil {
		w.Write(helper.ErrorHtml(templateErr.Error()))
		return
	}

	err := template.Execute(w, map[string]interface{}{
		"Title": "关于", "Body": string(markdown),
		"Config":config.Cfg,
	})
	if err != nil {
		w.Write(helper.ErrorHtml(err.Error()))
		return
	}

}
