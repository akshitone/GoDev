package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", prefixFunc)
	http.Handle("/mynamepath/", http.StripPrefix("/mynamepath", http.FileServer(http.Dir("./files/assets"))))
	http.ListenAndServe(":8080", nil)
}

func prefixFunc(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="/mynamepath/image.jpg" />`)
}
