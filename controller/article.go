package controller

import (
	"ForestBlog/models"
	"net/http"
)

func Article(w http.ResponseWriter, r *http.Request) {
	articleTemplate := models.Template.Article

	if err := r.ParseForm(); err != nil {
		articleTemplate.WriteError(w, err)
	}
	key := r.Form.Get("key")

	path := models.ArticleShortUrlMap[key]

	articleDetail, err := models.ReadArticleDetail(path)

	if err != nil {
		articleTemplate.WriteError(w, err)
	}

	articleTemplate.WriteData(w, models.BuildViewData("Article", articleDetail))
}
