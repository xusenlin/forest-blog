package models

import (
	"ForestBlog/config"
	"fmt"
	"html/template"
	"io"
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

	tp, err := readHtmlTemplate(
		[]string{"index", "extraNav", "dashboard", "categories", "article", "tags"},
		viewDir)
	if err != nil {
		return htmlTemplate, err
	}

	htmlTemplate.Index = tp[0]
	htmlTemplate.ExtraNav = tp[1]
	htmlTemplate.Dashboard = tp[2]
	htmlTemplate.Categories = tp[3]
	htmlTemplate.Article = tp[4]
	htmlTemplate.Tags = tp[5]

	return htmlTemplate, nil
}

func SpreadDigit(n int) []int {
	var r []int
	for i := 1; i <= n; i++ {
		r = append(r, i)
	}
	return r
}

func readHtmlTemplate(htmlFileName []string, viewDir string) ([]TemplatePointer, error) {
	var htmlTemplate []TemplatePointer

	head := viewDir + "/layouts/head.html"
	footer := viewDir + "/layouts/footer.html"

	for _, name := range htmlFileName {

		tp, err := template.New(name+".html").
			Funcs(template.FuncMap{"SpreadDigit": SpreadDigit}).
			ParseFiles(viewDir+"/"+name+".html", head, footer)
		if err != nil {
			return htmlTemplate, err
		}
		htmlTemplate = append(htmlTemplate, TemplatePointer{tp})
	}
	return htmlTemplate, nil
}
