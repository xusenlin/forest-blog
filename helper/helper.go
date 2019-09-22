package helper

import (
	"fmt"
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
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

func WriteErrorHtml(w http.ResponseWriter, err string) {
	_, newErr := w.Write(ErrorHtml(err))
	if newErr != nil {
		panic(newErr)
	}
}

func SedResponse(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write([]byte(`{"msg": "` + msg + `"}`))
	if err != nil {
		panic(err)
	}
}

func BuildArrByInt(num int) []int {
	var arr []int

	for i := 1; i <= num; i++ {
		arr = append(arr, i)
	}
	return arr
}

func UpdateArticle() {

	deleteCacheErr := os.RemoveAll("cache")
	if deleteCacheErr != nil {
		fmt.Println(deleteCacheErr)
	}

	blogPath := config.CurrentDir + "/" + config.Cfg.DocumentPath

	_, err := exec.LookPath("git")

	if err != nil {
		fmt.Println("请先安装git并克隆博客文档到" + blogPath)
		log.Fatalf("git cmd failed with %s\n", err)
	}

	cmd := exec.Command("git", "pull")
	cmd.Dir = blogPath

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	log.Println("UpdateArticle:" + string(out))
	//生成缓存
	_, err = models.GetMarkdownListByCache("/")

	if err != nil {
		log.Fatalf("生成缓存失败： %s\n", err)
	}
	return
}
