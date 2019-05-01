package controller

import (
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	type HomeInfo struct {
		Title string
		Articles []models.ArticleInfo
		Pagination []int
	}

	homeInfoData := HomeInfo{"首页",models.GetAllArticle(),[]int{1,2,3}}
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
