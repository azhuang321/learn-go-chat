package main

import (
	"chat/ctrl"
	"html/template"
	"net/http"
)

//注册视图
func RegisterView() {
	tpls, _ := template.ParseGlob("view/**/*")
	for _, v := range tpls.Templates() {
		tplName := v.Name()
		http.HandleFunc(tplName, func(w http.ResponseWriter, r *http.Request) {
			tpls.ExecuteTemplate(w, tplName, nil)
		})
	}
}

func main() {
	//绑定请求和处理函数
	http.HandleFunc("/user/login", ctrl.UserLogin)
	http.HandleFunc("/user/register", ctrl.UserRegister)

	http.HandleFunc("/contact/loadcommunity", ctrl.LoadCommunity)
	http.HandleFunc("/contact/loadfriend", ctrl.LoadFriend)
	http.HandleFunc("/contact/joincommunity", ctrl.JoinCommunity)
	//http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	http.HandleFunc("/contact/addfriend", ctrl.Addfriend)
	//http.HandleFunc("/chat", ctrl.Chat)
	//指定静态访问目录
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	RegisterView()
	//启动web
	http.ListenAndServe(":8080", nil)
}
