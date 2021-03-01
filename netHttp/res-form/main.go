package main

import (
	"fmt"
	"net/http"
)

type onetype int

func (one onetype) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Mcleod-Key", "GO Programming Form")
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(response, "<h1>Hi, I'm Akshit Mithaiwala.</h1>")
}

func main() {
	var a onetype
	http.ListenAndServe(":8080", a)
}
