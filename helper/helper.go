package helper

import "html/template"

func HtmlTemplate(fileName string) (* template.Template, error) {

	return template.ParseFiles(
		"resources/views/"+fileName+".html",
		"resources/views/layouts/head.html",
		"resources/views/layouts/footer.html")

}
