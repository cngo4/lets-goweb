package main

import (
	"net/http"
	"strings"
	"log"
	"html/template"
	"io"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	log.Println("Path => ", r.URL.Path)
	log.Println("Scheme => ", r.URL.Scheme)
	log.Println("url_long => ", r.Form["url_long"])
	for k, v := range r.Form {
		log.Println("key:", k)
		log.Println("val:", strings.Join(v, ""))
	}
	io.WriteString(w, "Hello golang!\n") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Println("request method is:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		if err := t.Execute(w, nil); err != nil {
			log.Fatal("login error", err)
		}
		log.Println("request:login.html")
	}
	if r.Method == "POST" {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		log.Println("username:", r.Form["username"])
		log.Println("password:", r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/", FormHandler)       //设置访问的路由
	http.HandleFunc("/login", login)         //设置访问的路由
	err := http.ListenAndServe(":12345", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}