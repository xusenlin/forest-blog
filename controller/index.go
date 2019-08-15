package controller

import (
	"fmt"
	"github.com/xusenlin/go_blog/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	_,err := models.GetMarkdownListByCache("/")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println("1")
	_,_ = w.Write([]byte("hell word"))
	return
	//fmt.Println(content)

	//_ = r.ParseForm()
	//
	//page,pageErr := strconv.Atoi(r.Form.Get("page"))
	//if pageErr != nil{
	//	page = 1
	//}
	//
	//template, templateErr := helper.HtmlTemplate("index")
	//if templateErr != nil {
	//	_,_ = w.Write(helper.ErrorHtml(templateErr.Error()))
	//	return
	//}

	//err := template.Execute(w, map[string]interface{}{
	//	"Title":"首页",
	//	"Data": models.GetMarkdownListByCache(page,""),
	//	"Config":config.Cfg,
	//})
	//
	//if err != nil {
	//	_,_ = w.Write(helper.ErrorHtml(err.Error()))
	//	return
	//}
}

//func Categories(w http.ResponseWriter, r *http.Request)  {
//
//	template, templateErr := helper.HtmlTemplate("categories")
//
//	if templateErr != nil {
//		_,_ = w.Write(helper.ErrorHtml(templateErr.Error()))
//		return
//	}
//
//	err := template.Execute(w, map[string]interface{}{
//		"Title":"分类",
//		"Data": models.GetCategoriesInfo(),
//		"Config":config.Cfg,
//	})
//	if err != nil {
//		_,_ = w.Write(helper.ErrorHtml(err.Error()))
//		return
//	}
//}
//
//func Works(w http.ResponseWriter, r *http.Request)  {
//
//	markdown,mdErr := models.GetMarkdownByPath("Works.md")
//	if mdErr != nil {
//		_,_ = w.Write(helper.ErrorHtml(mdErr.Error()))
//		return
//	}
//
//	template, templateErr := helper.HtmlTemplate("works")
//
//	if templateErr != nil {
//		_,_ = w.Write(helper.ErrorHtml(templateErr.Error()))
//		return
//	}
//
//	err := template.Execute(w, map[string]interface{}{
//		"Title": "作品",
//		"Body": string(markdown),
//		"Config":config.Cfg,
//	})
//	if err != nil {
//		_,_ = w.Write(helper.ErrorHtml(err.Error()))
//		return
//	}
//}
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
