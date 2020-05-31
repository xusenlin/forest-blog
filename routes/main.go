package routes

import (
	"ForestBlog/config"
	"ForestBlog/controller"
	"net/http"
)

func InitRoute()  {

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/blog", controller.Index)
	http.HandleFunc("/categories", controller.Category)
	http.HandleFunc("/article", controller.Article)
	http.HandleFunc("/extra-nav", controller.ExtraNav)

	http.HandleFunc(config.Cfg.GitHookUrl, controller.GithubHook)
	http.HandleFunc( config.Cfg.DashboardEntrance, controller.Dashboard)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(config.Cfg.DocumentAssetsDir))))

}