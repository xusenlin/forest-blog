package controller

import (
	"ForestBlog/models"
	"fmt"
	"net/http"
)

func GithubHook(w http.ResponseWriter, r *http.Request) {

	//err := r.ParseForm()
	//if err != nil {
	//	SedResponse(w, err.Error())
	//	return
	//}
	//
	//if "" == config.Cfg.WebHookSecret || "push" != r.Header.Get("x-github-event") {
	//	SedResponse(w, "No Configuration WebHookSecret Or Not Pushing Events")
	//	log.Println("No Configuration WebHookSecret Or Not Pushing Events")
	//	return
	//}
	//
	//sign := r.Header.Get("X-Hub-Signature")
	//
	//bodyContent, err := ioutil.ReadAll(r.Body)
	//
	//if err != nil {
	//	SedResponse(w, err.Error())
	//	log.Println("WebHook err:" + err.Error())
	//	return
	//}
	//
	//if err = r.Body.Close(); err != nil {
	//	SedResponse(w, err.Error())
	//	log.Println("WebHook err:" + err.Error())
	//	return
	//}
	//
	//mac := hmac.New(sha1.New, []byte(config.Cfg.WebHookSecret))
	//mac.Write(bodyContent)
	//expectedHash := "sha1=" + hex.EncodeToString(mac.Sum(nil))
	//
	//if sign != expectedHash {
	//	SedResponse(w, "WebHook err:Signature does not match")
	//	log.Println("WebHook err:Signature does not match")
	//	return
	//}

	SedResponse(w, "ok")

	models.CompiledContent()
}

func SedResponse(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write([]byte(`{"msg": "` + msg + `"}`))
	if err != nil {
		fmt.Println(err)
	}
}