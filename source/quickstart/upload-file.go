package main

import (
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"fmt"
	"os"
	"net/http"
	"html/template"
	"log"
)

// 处理/upload 逻辑
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("public/html/upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		// 解析文件
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		FilePath := "public/img/" + handler.Filename
		f, err := os.OpenFile(FilePath, os.O_WRONLY | os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		log.Println("file move to", FilePath)
	}
}

func main() {
	http.HandleFunc("/upload", UploadHandler)
	err := http.ListenAndServe(":12345", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}