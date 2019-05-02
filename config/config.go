package config

import (
	"encoding/json"
	"io/ioutil"
)

const ConfigFilePath = "app.json"


type Config struct {
	SiteName string `json:"siteName"`

	Port string `json:"port"`

	PageSize int `json:"pageSize"`

	DocumentPath string `json:"documentPath"`

	HtmlKeywords string `json:"htmlKeywords"`

	HtmlDescription string `json:"htmlDescription"`

	DocumentGithubUrl string `json:"documentGithubUrl"`

	MaxNumberArticleOfCategory int `json:"maxNumberArticleOfCategory"`
}

var Cfg  Config

func init()  {
	configFile,err := ioutil.ReadFile(ConfigFilePath)

	if err != nil {
		panic(err)
	}

	jsonErr := json.Unmarshal(configFile,&Cfg)

	if jsonErr != nil {
		panic(err)
	}

}