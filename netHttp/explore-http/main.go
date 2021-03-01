package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

type onetype int

func init() {
	tpl = template.Must(template.ParseFiles("netHttp/explore-http/index.gohtml"))
}

func (one onetype) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.ExecuteTemplate(writer, "index.gohtml", request.Form)
}

func main() {
	var a onetype
	http.ListenAndServe(":8080", a)
}
