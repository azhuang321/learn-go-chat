package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"html/template"
	"log"
	"net/http"
)

var DbEngin *xorm.Engine

func init() {
	driverName := "mysql"
	DsnName := "root:root@(127.0.0.1:3306)/chat?charset=utf8"
	DbEngin, err := xorm.NewEngine(driverName, DsnName)
	if err != nil {
		log.Fatal(err.Error())
	}
	DbEngin.ShowSQL(true)
	DbEngin.SetMaxOpenConns(2)

	fmt.Println("init database ok")
}

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
	//web 路由绑定
	http.HandleFunc("/user/login", userLogin)

	//指定静态访问目录
	http.Handle("/asset/", http.FileServer(http.Dir(".")))

	RegisterView()
	//启动web
	http.ListenAndServe(":8080", nil)
}
