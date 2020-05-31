package utils_test

import (
	"ForestBlog/utils"
	"testing"
)


func TestGetRepoName(t *testing.T){
	url := "https://github.com/xusenlin/ForestBlog.git"
	name,err := utils.GetRepoName(url)
	if  err != nil || name != "ForestBlog" {
		t.Error("repository name error")
	}
}
