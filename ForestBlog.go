package main

import (
	"ForestBlog/config"
	"ForestBlog/models"
	"ForestBlog/routes"
	"fmt"
	"net/http"
	"strconv"
)

func init() {
	models.CompiledContent()//克隆或者更新文章、递归生成文章、导航、短链 Map、加载模板
}

func main() {

	routes.InitRoute()
	fmt.Printf("Version：v%v \n" , config.Cfg.Version)
	fmt.Printf("ListenAndServe On Port %v \n" , config.Cfg.Port)
	fmt.Printf("UpdateArticle's GitHookUrl: %v   Secret:  %v \n" , config.Cfg.GitHookUrl,config.Cfg.WebHookSecret)
	if err := http.ListenAndServe( ":" + strconv.Itoa(config.Cfg.Port) , nil); err != nil{
		fmt.Println("ServeErr:",err)
	}
}
