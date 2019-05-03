package routes

import (
	"github.com/xusenlin/go_blog/controller"
	"net/http"
)

func InitRoute()  {

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/categories", controller.Categories)
	http.HandleFunc("/works", controller.Works)
	http.HandleFunc("/about", controller.About)
	//二级页面
	http.HandleFunc("/article", controller.Article)
	http.HandleFunc("/category", controller.CategoryArticle)
	//静态文件服务器
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("resources/public"))))
}