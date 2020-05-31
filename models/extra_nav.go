package models

import (
	"sort"
	"strings"
)

type Nav struct {
	Title string
	Path  string
}
type Navs []Nav

func initExtraNav(dir string) (Navs, error) {

	var navigation Navs
	var extraNav Articles

	extraNav, err := RecursiveReadArticles(dir)
	if err != nil {
		return navigation, err
	}
	sort.Sort(extraNav)

	for _, article := range extraNav {
		title := strings.Title(strings.ToLower(article.Title))
		navigation = append(navigation, Nav{Title: title, Path: article.Path})
	}

	return navigation, nil
}
