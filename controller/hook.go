package controller

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"github.com/xusenlin/go_blog/config"
	"github.com/xusenlin/go_blog/helper"
	"io/ioutil"
	"log"
	"net/http"
)

func GithubHook(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		helper.SedResponse(w, err.Error())
		return
	}

	if "" == config.Cfg.WebHookSecret || "push" != r.Header.Get("x-github-event") {
		helper.SedResponse(w, "No Configuration WebHookSecret Or Not Pushing Events")
		log.Println("No Configuration WebHookSecret Or Not Pushing Events")
		return
	}

	sign := r.Header.Get("X-Hub-Signature")

	bodyContent, err := ioutil.ReadAll(r.Body)

	if err != nil {
		helper.SedResponse(w, err.Error())
		log.Println("WebHook err:" + err.Error())
		return
	}

	if err = r.Body.Close(); err != nil {
		helper.SedResponse(w, err.Error())
		log.Println("WebHook err:" + err.Error())
		return
	}

	mac := hmac.New(sha1.New, []byte(config.Cfg.WebHookSecret))
	mac.Write(bodyContent)
	expectedHash := "sha1=" + hex.EncodeToString(mac.Sum(nil))

	if sign != expectedHash {
		helper.SedResponse(w, "WebHook err:Signature does not match")
		log.Println("WebHook err:Signature does not match")
		return
	}

	helper.SedResponse(w, "ok")

	helper.UpdateArticle()
}
