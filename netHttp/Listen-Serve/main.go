package main

import (
	"fmt"
	"net/http"
)

type onetype int

func (one onetype) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Hello! I'm Akshit Mithaiwala!")
}

func main() {
	var a onetype
	http.ListenAndServe(":8080", a)
}
