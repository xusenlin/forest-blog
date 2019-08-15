package models

import (
	"github.com/xusenlin/go_blog/config"
	"io/ioutil"
)

func GetCategories() (Categories, error) {

	var content Categories

	categoriesDir, err := ioutil.ReadDir(config.Cfg.DocumentPath + "/content")

	if err != nil {
		return content, err
	}

	for _, category := range categoriesDir {

		if !category.IsDir() {
			continue
		}
		var categoryContent Category
		markdownList, err := GetMarkdownListByCache(category.Name())

		if err != nil {
			return content, err
		}

		listLen := len(markdownList)
		categoryListFileNumber := listLen

		if listLen >= config.Cfg.CategoryListFileNumber {
			categoryListFileNumber = config.Cfg.CategoryListFileNumber
		}

		categoryContent.Title = category.Name()
		categoryContent.Number = listLen
		categoryContent.MarkdownFileList = markdownList[0:categoryListFileNumber]

		content = append(content, categoryContent)
	}

	return content, nil
}
