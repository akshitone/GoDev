package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("string-tpl/tpl-str.gohtml"))
}

func main() {
	str := []string{"zero", "one", "two"}
	data := struct {
		StrWord []string
		Name    string
	}{
		str, "Mithaiwala",
	}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
