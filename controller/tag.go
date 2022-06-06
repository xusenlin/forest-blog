package controller

import (
	"ForestBlog/config"
	"ForestBlog/models"
	"net/http"
)

func Tag(w http.ResponseWriter, r *http.Request)  {

	tagsTemplate := models.Template.Tags

	result := models.GroupByTag(&models.ArticleList,config.Cfg.TagDisplayQuantity)

	tagsTemplate.WriteData(w, models.BuildViewData("Tags",result))
}