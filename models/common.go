package models

import (
	"ForestBlog/config"
	"sync"
)

var Navigation Navs
var ArticleList Articles
var ArticleShortUrlMap map[string]string //用来保证文章 shortUrl 唯一和快速定位文章
var Template HtmlTemplate

func CompiledContent() {
	config.Initial() //克隆或者更新文档库
	//下面是对内容的生成
	wg := sync.WaitGroup{}
	var err error
	//导航
	wg.Add(1)
	go func() {
		Navigation, err = initExtraNav(config.Cfg.DocumentExtraNavDir)
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()

	//加载html模板
	wg.Add(1)
	go func() {
		Template, err = initHtmlTemplate(config.Cfg.CurrentDir + "/views")
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()

	//文章
	wg.Add(1)
	go func() {
		ArticleList, ArticleShortUrlMap, err = initArticles(config.Cfg.DocumentContentDir)
		if err != nil {
			panic(err)
		}
		wg.Done()
	}()
	wg.Wait()
	//启用并发比之前节约4倍左右的时间
	return
}
