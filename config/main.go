package config

import (
	"ForestBlog/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Config struct {
	userConfig
	systemConfig
}
//

var Cfg Config

func init() {
	var err error

	Cfg.CurrentDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	configFile, err := ioutil.ReadFile(Cfg.CurrentDir + "/config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(configFile, &Cfg)
	if err != nil {
		panic(err)
	}

	if "" == Cfg.DashboardEntrance ||
		! strings.HasPrefix(Cfg.DashboardEntrance, "/") {
		Cfg.DashboardEntrance = "/admin"
	}

	repoName, err := utils.GetRepoName(Cfg.DocumentGitUrl)

	if err != nil {
		panic(err)
	}

	Cfg.AppName = "ForestBlog"
	Cfg.Version = 3.0
	Cfg.DocumentDir = Cfg.CurrentDir + "/" + repoName
	Cfg.GitHookUrl = "/api/git_push_hook"
	Cfg.AppRepository = "https://github.com/xusenlin/ForestBlog"
}

func Initial() {
	if _, err := exec.LookPath("git"); err != nil {
		fmt.Println("请先安装git")
		panic(err)
	}
	if !utils.IsDir(Cfg.DocumentDir) {
		fmt.Println("正在克隆文档仓库，请稍等...")
		out, err := utils.RunCmdByDir(Cfg.CurrentDir, "git", "clone", Cfg.DocumentGitUrl)
		if err != nil {
			panic(err)
		}
		fmt.Println(out)
	} else {
		out, err := utils.RunCmdByDir(Cfg.DocumentDir, "git", "pull")
		fmt.Println(out)
		if err != nil {
			panic(err)
		}

	}
	if err := checkDocDirAndBindConfig(&Cfg); err != nil {
		fmt.Println("文档缺少必要的目录")
		panic(err)
	}

	imgDir := Cfg.CurrentDir + "/images"
	if !utils.IsDir(imgDir) {
		if os.Mkdir(imgDir, os.ModePerm) != nil {
			panic("生成images目录失败！")
		}
	}
}

func checkDocDirAndBindConfig(cfg *Config) error {
	dirs := []string{"assets", "content", "extra_nav"}
	for _, dir := range dirs {
		absoluteDir := Cfg.DocumentDir + "/" + dir
		if !utils.IsDir(absoluteDir) {
			return errors.New("documents cannot lack " + absoluteDir + " dir")
		}
	}
	cfg.DocumentAssetsDir = cfg.DocumentDir + "/assets"
	cfg.DocumentContentDir = cfg.DocumentDir + "/content"
	cfg.DocumentExtraNavDir = cfg.DocumentDir + "/extra_nav"
	return nil
}
