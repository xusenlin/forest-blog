package main

import (
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/routes"
	"net/http"
)

func main() {

	helper.StartTicker(func() {
		helper.UpdateArticle()
	})

	routes.InitRoute()

	http.ListenAndServe( ":" + config.Cfg.Port , nil)


}
