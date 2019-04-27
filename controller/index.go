package controller

import (
	"github.com/xusenlin/go_blog/helper"
	"github.com/xusenlin/go_blog/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	md ,_:= models.GetPostByPath()

	t, _ := helper.HtmlTemplate("index")
	err := t.Execute(w, map[string]string{"Title": "哈哈哈哈", "Body": string(md)})
	if err != nil {
		panic(err)
	}
}
