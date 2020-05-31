package utils_test

import (
	"ForestBlog/utils"
	"testing"
)


func TestGenerateShortUrl(t *testing.T){
	url := "https://github.com/xusenlin/ForestBlog.git"
	shortUrl := utils.GenerateShortUrl(url, func(url, keyword string) bool {return true})
	if  shortUrl != "ItMIz7" {
		t.Error("generate short URL error")
	}
}
