package helper

import (
	"bufio"
	"fmt"
	"github.com/xusenlin/go_blog/config"
	"html/template"
	"io"
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
	errorHtml :=`
			<div style='width: 100%;height: 100vh;display: flex;justify-content: center;align-items: center;'>
				<h1 style='padding: 10px 20px;background-color: #d9534f;color:#fff;border-radius: 4px;text-align: center;'
				onmouseover="this.style.backgroundColor='#f0ad4e';"
				>`+errorInfo+" :(</h1></div>"

	return []byte(errorHtml)
}
func BuildArrByInt(num int) []int {
	var arr []int

	for i := 1; i <= num; i ++ {
		arr = append(arr,i)
	}
	return arr
}

func StartTicker(f func()) {

	ticker:=time.NewTicker(time.Hour * 3)

	go func() {
		for _ =range ticker.C {
			f()
		}
	}()
}

func UpdateArticle()  {
	cmd := exec.Command("git","pull")
	//显示运行的命令
	cmd.Dir = config.CurrentDir + "/" + config.Cfg.DocumentPath

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}
	cmd.Wait()
	return
}

