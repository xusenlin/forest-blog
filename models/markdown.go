package models

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xusenlin/go_blog/config"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func GetMarkdown(path string) (Markdown, error) {
	//path=>categoryName/xxx.md
	fullPath := config.Cfg.DocumentPath + "/content" + path

	categoryName := strings.Replace(path,"/","",1)

	if strings.Index(categoryName,"/") == -1{ //文件在根目录下(content/)没有分类名称
		categoryName = ""
	}else {
		//虽然支持md文件无限层级，但是分类只算第一层目录，想做无限级分类的，但是考虑到每一级都可以放置md文件,分类页面不美观。
		categoryName = strings.Split(categoryName, "/")[0]
	}

	markdownFile, fileErr := os.Stat(fullPath)

	var content Markdown

	if fileErr != nil {
		return content, fileErr
	}
	if markdownFile.IsDir() {
		return content, errors.New("this path is Dir")
	}
	markdown, mdErr := ioutil.ReadFile(fullPath)

	if mdErr != nil {
		return content, mdErr
	}
	markdown = bytes.TrimSpace(markdown)

	content.Category = categoryName
	content.Title = markdownFile.Name()
	content.Date = Time(markdownFile.ModTime())
	content.Path = path

	desc := []rune(string(markdown))
	descLen := len(desc)
	if descLen <= 150 {
		content.Description = string(desc[0:descLen])
	}else {
		content.Description = string(desc[0:150])
	}

	if ! bytes.HasPrefix(markdown, []byte("```json")) {
		return content, nil
	}

	markdown = bytes.Replace(markdown, []byte("```json"), []byte(""), 1)
	markdownInfo := bytes.SplitN(markdown, []byte("```"), 2)[0]

	if err := json.Unmarshal(bytes.TrimSpace(markdownInfo), &content); err != nil {
		return content, err
	}

	return content, nil
}

func GetMarkdownDetails(path string) (MarkdownDetails, error) {
	//path=>categoryName/xxx.md
	fullPath := config.Cfg.DocumentPath + "/content/" + path

	categoryName := strings.Replace(path,"/","",1)

	if strings.Index(categoryName,"/") == -1{ //文件在根目录下(content/)没有分类名称
		categoryName = ""
	}else {
		categoryName = strings.Split(categoryName, "/")[0]
	}

	markdownFile, fileErr := os.Stat(fullPath)

	var content MarkdownDetails

	if fileErr != nil {
		return content, fileErr
	}
	if markdownFile.IsDir() {
		return content, errors.New("this path is Dir")
	}
	markdown, mdErr := ioutil.ReadFile(fullPath)

	if mdErr != nil {
		return content, mdErr
	}
	markdown = bytes.TrimSpace(markdown)

	content.Path = path
	content.Body = string(markdown)
	content.Category = categoryName
	content.Title = markdownFile.Name()
	content.Date = Time(markdownFile.ModTime())

	if ! bytes.HasPrefix(markdown, []byte("```json")) {
		return content, nil
	}

	markdown = bytes.Replace(markdown, []byte("```json"), []byte(""), 1)
	markdownArrInfo := bytes.SplitN(markdown, []byte("```"), 2)

	content.Body = string(markdownArrInfo[1])

	if err := json.Unmarshal(bytes.TrimSpace(markdownArrInfo[0]), &content); err != nil {
		return content, err
	}

	return content, nil
}

func GetMarkdownList(dir string) (MarkdownList, error) {
	//path=>categoryName
	var fullDir string
	fullDir = config.Cfg.DocumentPath + "/content" + dir

	fileOrDir, err := ioutil.ReadDir(fullDir)

	var mdList MarkdownList

	if err != nil {
		return mdList, err
	}

	for _, fileInfo := range fileOrDir {
		var subDir string
		if "/" == dir {
			subDir = "/" + fileInfo.Name()
		} else {
			subDir = dir + "/" + fileInfo.Name()
		}
		if fileInfo.IsDir() {

			subMdList, err := GetMarkdownList(subDir)
			if err != nil {
				return mdList, err
			}
			mdList = append(mdList, subMdList...)
		} else if strings.HasSuffix(strings.ToLower(fileInfo.Name()), "md") {
			markdown, err := GetMarkdown(subDir)
			if err != nil {
				return mdList, err
			}
			mdList = append(mdList, markdown)
		}
	}
	return mdList, nil

}

func GetMarkdownListByCache(dir string) (MarkdownList, error) {

	cacheFileName := fmt.Sprintf("%x", md5.Sum([]byte(dir)))

	cacheFilePath := config.CurrentDir + "/cache/" + cacheFileName + ".json"

	var content MarkdownList

	cacheFile, cacheErr := ioutil.ReadFile(cacheFilePath)

	if cacheErr == nil && json.Unmarshal(cacheFile, &content) == nil {
		return content, nil
	}
	fmt.Println("没有缓存！！")
	content, err := GetMarkdownList(dir)

	if err != nil {
		return content, err
	}

	sort.Sort(content)
	markdownListJson, err := json.Marshal(content)

	if err != nil {
		return content, err
	}

	err = ioutil.WriteFile(cacheFilePath, markdownListJson, os.ModePerm)

	if err != nil {
		return content, err
	}

	return content, nil
}
