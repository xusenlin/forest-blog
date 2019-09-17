package controller

import (
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/service"
	"net/http"
	"strconv"
)

func Dashboard(w http.ResponseWriter, r *http.Request)  {

	var dashboardMsg []string

	err := r.ParseForm()
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	index, err := strconv.Atoi(r.Form.Get("theme"))
	if err == nil && index < len(config.Cfg.ThemeOption) {
		service.SetThemeColor(index)
		dashboardMsg = append(dashboardMsg,"颜色切换成功!")
	}

	action := r.Form.Get("action")
	if "updateArticle" == action {
		helper.UpdateArticle()
		dashboardMsg = append(dashboardMsg,"文章更新成功!")
	}

	template, err := helper.HtmlTemplate("dashboard")

	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "控制台",
		"Data":   map[string]interface{}{
			"msg":dashboardMsg,
		},
		"Config": config.Cfg,
	})
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

}