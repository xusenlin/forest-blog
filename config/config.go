package config

const configFilePath = "./config.json"


type Config struct {
	Name string `json:"siteName"`

	Port uint `json:"port"`

	DocumentGithubUrl string `json:"documentGithubUrl"`
}

