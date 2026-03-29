package models

import (
	"ForestBlog/config"
	"fmt"
	"io"
	"text/template"
)

type TemplatePointer struct {
	*template.Template
}

type HtmlTemplate struct {
	Article    TemplatePointer
	Categories TemplatePointer
	Tags       TemplatePointer
	Dashboard  TemplatePointer
	ExtraNav   TemplatePointer
	Index      TemplatePointer
}

func (t TemplatePointer) WriteData(w io.Writer, data interface{}) {

	err := t.Execute(w, data)
	if err != nil {
		if _, e := w.Write([]byte(err.Error())); e != nil {
			fmt.Println(e)
		}
	}
}

func (t TemplatePointer) WriteError(w io.Writer, err error) {
	if _, e := w.Write([]byte(err.Error())); e != nil {
		fmt.Println(e)
	}
}

func BuildViewData(title string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Title":  title,
		"Data":   data,
		"Config": config.Cfg,
		"Navs":   Navigation,
	}
}

func initHtmlTemplate(viewDir string) (HtmlTemplate, error) {
	var htmlTemplate HtmlTemplate

	var err error

	if htmlTemplate.Index, err = readHtmlTemplate("index", viewDir); err != nil {
		return htmlTemplate, err
	}
	if htmlTemplate.ExtraNav, err = readHtmlTemplate("extraNav", viewDir); err != nil {
		return htmlTemplate, err
	}
	if htmlTemplate.Dashboard, err = readHtmlTemplate("dashboard", viewDir); err != nil {
		return htmlTemplate, err
	}
	if htmlTemplate.Categories, err = readHtmlTemplate("categories", viewDir); err != nil {
		return htmlTemplate, err
	}
	if htmlTemplate.Article, err = readHtmlTemplate("article", viewDir); err != nil {
		return htmlTemplate, err
	}
	if htmlTemplate.Tags, err = readHtmlTemplate("tags", viewDir); err != nil {
		return htmlTemplate, err
	}

	return htmlTemplate, nil
}

func SpreadDigit(currentPage, totalPage int) []int {
	// 如果总页数小于等于7，显示所有页码
	if totalPage <= 7 {
		var r []int
		for i := 1; i <= totalPage; i++ {
			r = append(r, i)
		}
		return r
	}

	var r []int
	// 总是显示第一页
	r = append(r, 1)

	// 计算当前页附近的页码范围
	start := currentPage - 2
	end := currentPage + 2

	// 调整范围，确保不超出边界
	if start < 2 {
		start = 2
		end = start + 4
	}
	if end > totalPage-1 {
		end = totalPage - 1
		start = end - 4
	}

	// 如果需要，在第一页和当前页范围之间添加省略号
	if start > 2 {
		r = append(r, 0) // 0 表示省略号
	}

	// 添加当前页附近的页码
	for i := start; i <= end; i++ {
		r = append(r, i)
	}

	// 如果需要，在当前页范围和最后一页之间添加省略号
	if end < totalPage-1 {
		r = append(r, 0) // 0 表示省略号
	}

	// 总是显示最后一页
	r = append(r, totalPage)

	return r
}

func readHtmlTemplate(htmlFileName string, viewDir string) (TemplatePointer, error) {

	head := viewDir + "/layouts/head.html"
	footer := viewDir + "/layouts/footer.html"

	tp, err := template.New(htmlFileName+".html").
		Funcs(template.FuncMap{"SpreadDigit": SpreadDigit}).
		ParseFiles(viewDir+"/"+htmlFileName+".html", head, footer)
	if err != nil {
		return TemplatePointer{}, err
	}
	return TemplatePointer{tp}, nil
}
