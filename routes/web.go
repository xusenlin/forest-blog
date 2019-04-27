package routes

import (
	"github.com/xusenlin/go_blog/controller"
	"net/http"
)

func InitRoute()  {

	http.HandleFunc("/", controller.Index)

	//http.Handle("/pollux/", http.StripPrefix("/pollux/", http.FileServer(http.Dir("file"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("resources/public"))))
	//fsh := http.FileServer(http.Dir("resources/public"))
	//http.Handle("/static/", http.StripPrefix("/static/", fsh))
}