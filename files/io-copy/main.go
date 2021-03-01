package main

import (
	"io"
	"net/http"
)

func fileFunc(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<img src="https://avatars.githubusercontent.com/u/58538754?s=460&u=b7d57aca8b442f936e89227e02dbf615c1b3094f&v=4" />
	`)
}

func main() {
	http.HandleFunc("/", fileFunc)
	http.ListenAndServe(":8080", nil)
}
