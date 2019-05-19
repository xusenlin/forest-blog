package main

import (
	"fmt"
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

	fmt.Println("服务已经启动 Listening...")

	http.ListenAndServe( ":" + config.Cfg.Port , nil)

}
