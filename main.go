package main

import (
	"net/http"
	"github.com/xusenlin/go_blog/routes"
)

func main() {
	routes.InitRoute()
	http.ListenAndServe(":80", nil)
}
