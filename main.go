package main

import (
	"fmt"
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/routes"
	"net/http"
	"strconv"
)

func main() {

	routes.InitRoute()

	fmt.Println(config.Cfg.AppName)
	fmt.Printf("Version：v%v \n" , config.Cfg.Version)
	fmt.Printf("ListenAndServe On Port %v \n" , config.Cfg.Port)
	fmt.Printf("UpdateArticle's GitHookUrl: %v   Secret:  %v \n" , config.Cfg.GitHookUrl,config.Cfg.WebHookSecret)

	helper.UpdateArticle()

	if err := http.ListenAndServe( ":" + strconv.Itoa(config.Cfg.Port) , nil); err != nil{
		fmt.Println("ServeErr:",err)
	}

}
