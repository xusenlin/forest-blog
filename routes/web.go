package routes

import (
	"github.com/xusenlin/go_blog/controller"
	"net/http"
)

func InitRoute()  {

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/blog", controller.Index)
	http.HandleFunc("/categories", controller.Categories)
	http.HandleFunc("/works", controller.Works)
	http.HandleFunc("/about", controller.About)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("resources/public"))))
}