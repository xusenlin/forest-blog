package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const ConfigFilePath = "app.json"

type Config struct {
	SiteName string `json:"siteName"`

	Author string `json:"author"`

	Icp string `json:"icp"`

	TimeLayout string `json:"timeLayout"`

	Port string `json:"port"`

	UpdateArticleInterval int `json:"updateArticleInterval"`

	PageSize int `json:"pageSize"`

	DescriptionLen int `json:"descriptionLen"`

	DocumentPath string `json:"documentPath"`

	HtmlKeywords string `json:"htmlKeywords"`

	HtmlDescription string `json:"htmlDescription"`

	CategoryListFileNumber int `json:"categoryListFileNumber"`

	ThemeColor string `json:"themeColor"`

	ThemeOption []string `json:"themeOption"`

	DashboardEntrance string `json:"dashboardEntrance"`
}

var Cfg Config

var CurrentDir string

func init() {
	var pwdErr error

	CurrentDir, pwdErr = os.Getwd()

	if pwdErr != nil {
		panic(pwdErr)
	}

	configFile, err := ioutil.ReadFile(ConfigFilePath)

	if err != nil {
		panic(err)
	}

	jsonErr := json.Unmarshal(configFile, &Cfg)

	if jsonErr != nil {
		panic(err)
	}

	fmt.Println("init config...")
}
