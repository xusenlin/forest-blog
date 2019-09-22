package routes

import (
	"github.com/xusenlin/go_blog/controller"
	"net/http"
)

func initApiRoute()  {

	http.HandleFunc("/api/git_push_hook", controller.GithubHook)

}