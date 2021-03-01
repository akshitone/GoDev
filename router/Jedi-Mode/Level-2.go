package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("router/Jedi-Mode/Level-2.gohtml"))
}

func indexFunc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi world, It's index page!")
}
func dogFunc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Hi world, It's dog page!")
}
func meFunc(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "Level-2.gohtml", "Akshit Mithaiwala")
	if err != nil {
		log.Fatalln("error executing template:", err)
	}
}

func main() {
	http.HandleFunc("/", indexFunc)
	http.HandleFunc("/dog/", dogFunc)
	http.HandleFunc("/me/", meFunc)

	http.ListenAndServe(":8080", nil)
}
