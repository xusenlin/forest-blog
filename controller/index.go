package controller

import (
	"fmt"
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {


	template, templateErr := helper.HtmlTemplate("index")
	if templateErr != nil {
		panic(templateErr)
	}

	err := template.Execute(w, map[string]string{"Title": "首页", "Body": "阿斯顿撒多"})
	if err != nil {
		panic(err)
	}
}

func Categories(w http.ResponseWriter, r *http.Request)  {

	fmt.Println(models.GetCategoriesInfo())

	template, templateErr := helper.HtmlTemplate("categories")

	if templateErr != nil {
		panic(templateErr)
	}

	err := template.Execute(w, map[string]string{"Title": "分类", "Body": "阿斯顿撒多"})
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
