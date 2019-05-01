package models

import (
	"github.com/xusenlin/go_blog/config"
	"io/ioutil"
)

func GetMarkdownByPath(documentPath string) ([]byte, error) {
	return ioutil.ReadFile(config.Cfg.DocumentPath + "/" + documentPath)
}
