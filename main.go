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
		helper.ClearCache()
	})

	routes.InitRoute()

	fmt.Println("Listening...")

	if err := http.ListenAndServe( ":" + config.Cfg.Port , nil); err != nil{
		fmt.Println("ServeErr:",err)
	}
}
