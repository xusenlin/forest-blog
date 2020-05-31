package controller

import (
	"ForestBlog/config"
	"ForestBlog/models"
	"net/http"
)

func Category(w http.ResponseWriter, r *http.Request)  {

	categoriesTemplate := models.Template.Categories

	result := models.GroupByCategory(&models.ArticleList,config.Cfg.CategoryDisplayQuantity)

	categoriesTemplate.WriteData(w, models.BuildViewData("Blog",result))
}