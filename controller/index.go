package controller

import (
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/models"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	type HomeInfo struct {
		Title string
		Pagination models.ArticlesPagination
	}

	page,pageErr := strconv.Atoi(r.Form.Get("page"))
	if pageErr != nil{
		page = 1
	}
	homeInfoData := HomeInfo{"首页",models.GetArticleByPage(page)}

	template, templateErr := helper.HtmlTemplate("index")
	if templateErr != nil {
		panic(templateErr)
	}

	err := template.Execute(w, map[string]HomeInfo{"Data": homeInfoData})
	if err != nil {
		panic(err)
	}
}

func Categories(w http.ResponseWriter, r *http.Request)  {

	type CategoryInfo struct {
		Title string
		Categories []models.Category
	}
	categoryInfoData := CategoryInfo{"分类",models.GetCategoriesInfo()}

	template, templateErr := helper.HtmlTemplate("categories")

	if templateErr != nil {
		panic(templateErr)
	}

	err := template.Execute(w, map[string]CategoryInfo{"Data": categoryInfoData})
	if err != nil {
		panic(err)
	}
}
func Works(w http.ResponseWriter, r *http.Request)  {

	markdown,mdErr := models.GetMarkdownByPath("Works.md")
	if mdErr != nil {
		panic(mdErr)
	}

	template, templateErr := helper.HtmlTemplate("works")

	if templateErr != nil {
		panic(templateErr)
	}

	err := template.Execute(w, map[string]string{"Title": "作品", "Body": string(markdown)})
	if err != nil {
		panic(err)
	}
}

func About(w http.ResponseWriter, r *http.Request)  {

	markdown,mdErr := models.GetMarkdownByPath("About.md")
	if mdErr != nil {
		panic(mdErr)
	}

	template, templateErr := helper.HtmlTemplate("about")
	if templateErr != nil {
		panic(templateErr)
	}

	err := template.Execute(w, map[string]string{"Title": "关于", "Body": string(markdown)})
	if err != nil {
		panic(err)
	}

}
