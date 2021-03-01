package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

var tpl *template.Template

type onetype int

func init() {
	tpl = template.Must(template.ParseFiles("netHttp/methods-urls/index.gohtml"))
}

func (one onetype) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions map[string][]string
		Header      http.Header
	}{
		request.Method,
		request.URL,
		request.Form,
		request.Header,
	}

	tpl.ExecuteTemplate(writer, "index.gohtml", data)
}

func main() {
	var a onetype
	http.ListenAndServe(":8080", a)
}
