package controller

import (
	"ForestBlog/models"
	"net/http"
)

func ExtraNav(w http.ResponseWriter, r *http.Request) {
	extraNavTemplate := models.Template.ExtraNav

	if err := r.ParseForm(); err != nil {
		extraNavTemplate.WriteError(w, err)
	}


	name := r.Form.Get("name")
	for _,nav := range models.Navigation {
		if nav.Title == name {
			articleDetail, err := models.ReadArticleDetail(nav.Path)
			if err != nil {
				extraNavTemplate.WriteError(w, err)
			}
			extraNavTemplate.WriteData(w, models.BuildViewData(nav.Title,articleDetail))
			return
		}
	}
}