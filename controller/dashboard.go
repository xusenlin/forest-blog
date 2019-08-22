package controller

import (
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/service"
	"net/http"
	"strconv"
)

func Dashboard(w http.ResponseWriter, r *http.Request)  {
	err := r.ParseForm()
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	index, err := strconv.Atoi(r.Form.Get("color_index"))
	if err == nil && index < len(config.Cfg.ThemeOption) {
		service.SetThemeColor(index)
	}

	template, err := helper.HtmlTemplate("dashboard")

	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

	err = template.Execute(w, map[string]interface{}{
		"Title":  "控制台",
		"Data":   "",
		"Config": config.Cfg,
	})
	if err != nil {
		helper.WriteErrorHtml(w, err.Error())
		return
	}

}