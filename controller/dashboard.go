package controller

import (
	"ForestBlog/config"
	"ForestBlog/models"
	"net/http"
	"strconv"
)

func Dashboard(w http.ResponseWriter, r *http.Request)  {

	var dashboardMsg []string
	dashboardTemplate := models.Template.Dashboard

	if err := r.ParseForm(); err != nil {
		dashboardTemplate.WriteError(w, err)
	}

	index, err := strconv.Atoi(r.Form.Get("theme"))
	if err == nil && index < len(config.Cfg.ThemeOption) {
		config.Cfg.ThemeColor = config.Cfg.ThemeOption[index]
		dashboardMsg = append(dashboardMsg,"颜色切换成功!")
	}

	action := r.Form.Get("action")
	if "updateArticle" == action {
		models.CompiledContent()
		dashboardMsg = append(dashboardMsg,"文章更新成功!")
	}

	dashboardTemplate.WriteData(w, models.BuildViewData("Dashboard", map[string]interface{}{
		"msg":dashboardMsg,
	}))

}
