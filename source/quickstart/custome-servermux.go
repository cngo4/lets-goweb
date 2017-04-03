package main

import (
	"net/http"
	"io"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		MyHelloServer(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func MyHelloServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, ServerMux!\n")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":12345", mux)
}
