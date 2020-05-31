package utils

import (
	"errors"
	"strings"
)

//通过git url 返回仓库的名字
func GetRepoName(gitUrl string) (string, error) {

	if !strings.HasSuffix(gitUrl, ".git") {
		return "", errors.New("git URL must end with .git！")
	}

	noSuffixUrl := strings.TrimSuffix(gitUrl, ".git")
	urlArr := strings.Split(noSuffixUrl, "/")

	return urlArr[ len(urlArr)-1 ], nil
}
