package main

import (
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/routes"
	"net/http"
)

func main() {

	routes.InitRoute()

	http.ListenAndServe( ":" + config.Cfg.Port , nil)
}
