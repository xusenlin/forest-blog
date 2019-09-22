package main

import (
	"fmt"
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/routes"
	"net/http"
)

func main() {

	routes.InitRoute()

	fmt.Println(config.Cfg.AppName)
	fmt.Printf("Version：v%v \n" , config.Cfg.Version)

	helper.UpdateArticle()

	if err := http.ListenAndServe( ":" + config.Cfg.Port , nil); err != nil{
		fmt.Println("ServeErr:",err)
	}

}
