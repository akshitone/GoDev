package main

import (
	"io"
	"net/http"
)

func indexFunc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi world, It's index page!")
}
func dogFunc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi world, It's dog page!")
}
func meFunc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi Akshit Mithaiwala!")
}

func main() {
	http.HandleFunc("/", indexFunc)
	http.HandleFunc("/dog/", dogFunc)
	http.HandleFunc("/me/", meFunc)

	http.ListenAndServe(":8080", nil)
}
