package main

import (
	"io"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("files/Jedi-Mode/Level-1.gohtml"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", fooFunc)
	http.HandleFunc("/dog/", dogFunc)
	http.HandleFunc("/dogimage", dogImageFunc)

	http.ListenAndServe(":8080", nil)
}

func fooFunc(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, "Foo Run!")
}

func dogFunc(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "Level-1.gohtml", "Akshit Mithaiwala")
}

func dogImageFunc(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "files/Jedi-Mode/dog.jpg")
}
