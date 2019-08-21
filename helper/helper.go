package helper

import (
	"fmt"
	"github.com/xusenlin/go_blog/config"
	"html/template"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func HtmlTemplate(fileName string) (*template.Template, error) {

	return template.ParseFiles(
		"resources/views/"+fileName+".html",
		"resources/views/layouts/head.html",
		"resources/views/layouts/footer.html")
}

func ErrorHtml(errorInfo string) []byte {
	errorHtml := `
			<div style='width: 100%;height: 100vh;display: flex;justify-content: center;align-items: center;'>
				<p style='padding: 10px 20px;background-color: #d9534f;color:#fff;border-radius: 4px;text-align: center;'
				onmouseover="this.style.backgroundColor='#f0ad4e';"
				>` + errorInfo + " :(</p></div>"

	return []byte(errorHtml)
}

func WriteErrorHtml(w http.ResponseWriter,err string)  {
	_,newErr := w.Write(ErrorHtml(err))
	if newErr != nil{
		panic(newErr)
	}
}

func BuildArrByInt(num int) []int {
	var arr []int

	for i := 1; i <= num; i++ {
		arr = append(arr, i)
	}
	return arr
}

func StartTicker(f func()) {

	updateTime := time.Duration(config.Cfg.UpdateArticleInterval)

	ticker := time.NewTicker(time.Hour * updateTime)

	go func() {
		for _ = range ticker.C {
			f()
		}
	}()
}

func UpdateArticle() {

	cmd := exec.Command("git", "pull")
	cmd.Dir = config.CurrentDir + "/" + config.Cfg.DocumentPath

	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	return
}

func ClearCache()  {
	err := os.RemoveAll("cache")
	if err != nil{
		fmt.Println(err)
		return
	}
}
