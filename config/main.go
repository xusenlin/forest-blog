package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	userConfig
	systemConfig
}


var Cfg Config
var CurrentDir string

func init() {
	var pwdErr error

	CurrentDir, pwdErr = os.Getwd()

	if pwdErr != nil {
		panic(pwdErr)
	}

	configFile, err := ioutil.ReadFile("config.json")

	if err != nil {
		panic(err)
	}

	jsonErr := json.Unmarshal(configFile, &Cfg)

	if jsonErr != nil {
		panic(err)
	}
	if "" == Cfg.DashboardEntrance || ! strings.HasPrefix(Cfg.DashboardEntrance, "/") {
		Cfg.DashboardEntrance = "/admin"
	}

	Cfg.AppName = "ForestBlog"
	Cfg.Version = 2.2
	Cfg.GitHookUrl = "/api/git_push_hook"
	Cfg.AppRepository = "https://github.com/xusenlin/ForestBlog"

}
