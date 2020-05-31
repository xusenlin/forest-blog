package utils_test

import (
	"ForestBlog/utils"
	"testing"
)


func TestRunCmdByDir(t *testing.T){
	_,err := utils.RunCmdByDir("./","ping","127.0.0.1")
	if  err != nil {
		t.Error("run cmd error",err)
	}
}
