package models

import "io/ioutil"

func GetPostByPath() ([]byte, error) {
	return ioutil.ReadFile("resources/blog_docs/About.md")
}
