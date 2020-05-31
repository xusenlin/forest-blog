package controller

import (
	"ForestBlog/config"
	"ForestBlog/models"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {

	indexTemplate := models.Template.Index

	if err := r.ParseForm();err != nil {
		indexTemplate.WriteError(w, err)
	}

	page, err := strconv.Atoi(r.Form.Get("page"))
	if err != nil {
		page = 1
	}
	articles := models.ArticleList

	search := r.Form.Get("search")
	category := r.Form.Get("category")

	if  search != "" || category != ""{
		articles = models.ArticleSearch(&articles,search,category)
	}

	result := models.Pagination(&articles,page,config.Cfg.PageSize)

	indexTemplate.WriteData(w, models.BuildViewData("Blog",result))
}


