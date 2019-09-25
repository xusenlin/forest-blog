package routes

import (
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/controller"
	"net/http"
)

func initApiRoute()  {

	http.HandleFunc(config.Cfg.GitHookUrl, controller.GithubHook)

}