package main

import (
	"encoding/json"
	"net/http"
)

type H struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg"`
}

func Resp(w http.ResponseWriter, Code int, Data interface{}, Msg string) {
	h := H{
		Code: Code,
		Data: Data,
		Msg:  Msg,
	}
	str, _ := json.Marshal(h)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(str))
}

func userLogin(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	mobile := request.PostForm.Get("mobile")
	password := request.PostForm.Get("password")
	loginOk := false

	if mobile == "18612345678" && password == "123123" {
		loginOk = true
	}
	if loginOk {
		Resp(writer, 0, map[string]interface{}{
			"id":    1,
			"token": "test",
		}, "")
	} else {
		Resp(writer, -1, nil, "密码不正确")
	}

}

func main() {
	//web 路由绑定
	http.HandleFunc("/user/login", userLogin)

	//指定静态访问目录
	http.Handle("/asset/", http.FileServer(http.Dir(".")))
	//启动web
	http.ListenAndServe(":8080", nil)
}
